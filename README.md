# Gataca OIDC Test Client (Go)

This project demonstrates a simple OpenID Connect (OIDC) client built in Go, using [Gataca's Vouch OIDC provider](https://vouch.gataca.io). It uses the Authorization Code flow to authenticate a user and validate the resulting ID token.

For this example, I used the scope: over18fae. To use other scopes, it's necessary to check the Gataca Studio configuration.

---

## 🔧 Configuration

Configuration is loaded via a `config.yaml` file using the [`viper`](https://github.com/spf13/viper) library.

Create a `config.yaml` file in the root directory:

```yaml
client_id: YOUR_CLIENT_ID
client_secret: YOUR_CLIENT_SECRET
redirect_uri: https://yourdomain.com/callback

