package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

var (
	clientID     string
	clientSecret string
	redirectURL  string
	providerURL  = "https://vouch.gataca.io/"

	oauth2Config *oauth2.Config
	verifier     *oidc.IDTokenVerifier
)

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	clientID = viper.GetString("client_id")
	clientSecret = viper.GetString("client_secret")
	redirectURL = viper.GetString("redirect_uri")

	if clientID == "" || clientSecret == "" || redirectURL == "" {
		log.Fatal("Missing required configuration in .config.yaml")
	}
}

func main() {
	initConfig()

	ctx := context.Background()

	provider, err := oidc.NewProvider(ctx, providerURL)
	if err != nil {
		log.Fatalf("Failed to get provider: %v", err)
	}

	oauth2Config = &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Endpoint:     provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "over18fae"},
	}

	verifier = provider.Verifier(&oidc.Config{ClientID: clientID})

	http.HandleFunc("/", handleLogin)
	http.HandleFunc("/callback", handleCallback)

	fmt.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	state := "random-state"
	authCodeURL := oauth2Config.AuthCodeURL(state)
	http.Redirect(w, r, authCodeURL, http.StatusFound)
}

func handleCallback(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	code := r.URL.Query().Get("code")

	token, err := oauth2Config.Exchange(ctx, code)
	if err != nil {
		http.Error(w, "Token exchange failed", http.StatusInternalServerError)
		log.Println("Token exchange error:", err)
		return
	}

	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		http.Error(w, "No id_token in token response", http.StatusInternalServerError)
		return
	}

	idToken, err := verifier.Verify(ctx, rawIDToken)
	if err != nil {
		http.Error(w, "Failed to verify ID Token", http.StatusInternalServerError)
		log.Println("ID Token verification error:", err)
		return
	}

	var claims map[string]interface{}
	if err := idToken.Claims(&claims); err != nil {
		http.Error(w, "Failed to parse claims", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Login successful!\n\n")
	fmt.Fprintf(w, "Access Token:\n%s\n\n", token.AccessToken)
	fmt.Fprintf(w, "ID Token:\n%s\n\n", rawIDToken)
	fmt.Fprintf(w, "ID Token Claims:\n%+v\n", claims)
}
