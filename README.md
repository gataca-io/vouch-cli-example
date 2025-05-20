# Gataca OIDC Test Client (Go)

This repository showcases a minimal OpenID Connect (OIDC) client built in Go, using [Gataca's Vouch OIDC provider](https://vouch.gataca.io). It follows the Authorization Code flow to authenticate a user and validate the received ID token.

> This example uses the `over18fae` scope. If you want to use different scopes, make sure they are enabled in your **Gataca Studio** App Integration settings and listed in your local configuration.

---

## 🚀 Local Deployment

OIDC clients typically require exact domain matching for redirect URIs, which can make local development tricky. To simulate a valid domain for redirects, this setup uses [**Ngrok**](https://ngrok.com/), exposing a local server running on port `:8080` to a public `.ngrok-free.app` domain.

### 1. Start Ngrok

```bash
ngrok http 8080
```

This creates a public URL (e.g., `https://random-string.ngrok-free.app`) that will forward traffic to your local port 8080.

### 2. Update Configuration

Edit your `config.yaml` and replace the redirect URI with the Ngrok URL (+ `/callback`):

```yaml
redirect_uri: https://random-string.ngrok-free.app/callback
```

### 3. Run the OIDC Client

```bash
go run main.go
```

### 4. Update Vouch Redirect URI

In Gataca Studio, update the **App Integration** (e.g., "Vouch - IdP") to include the Ngrok `redirect_uri` under the list of allowed values.

### 5. Start the Flow

Open the Ngrok URL (without `/callback`) in your browser. This starts the OIDC login flow. After authentication, you’ll be redirected back and your ID token will be validated.

---

## 🔧 Configuration

This project uses [`viper`](https://github.com/spf13/viper) for config management.

Create a file named `config.yaml` in the root directory with the following content:

```yaml
client_id: YOUR_CLIENT_ID
client_secret: YOUR_CLIENT_SECRET
redirect_uri: https://random-string.ngrok-free.app/callback
```

---

## 📂 Structure

- `main.go`: Entry point of the app, handles login and callback endpoints.
- `config.yaml`: Local configuration (not committed).
- `.gitignore`: Includes standard Go ignores and excludes `config.yaml`.

---

## 📝 Notes

- Only scopes listed in the App Integration configuration are valid. Attempting to use others will result in errors.
- You can pre-configure multiple scopes in Gataca Studio, and selectively use a subset during implementation.
- FAE activation is transparent to the OIDC flow; it does not affect the integration logic.

---

## 🛠️ Dependencies

- [`golang.org/x/oauth2`](https://pkg.go.dev/golang.org/x/oauth2)
- [`github.com/coreos/go-oidc`](https://github.com/coreos/go-oidc)
- [`github.com/spf13/viper`](https://github.com/spf13/viper)

---

## 📫 Contact

If you have any questions or issues, feel free to open an issue or reach out via [gataca.io](https://gataca.io).
