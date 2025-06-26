📘 README: Golang OIDC + Okta Auth Web App

This is a small Golang web application that integrates OIDC (OpenID Connect) authentication using Okta as the identity provider. It allows users to log in via Okta and retrieves their profile information using ID tokens.

%% Features
✅ Login with Okta via OIDC
✅ Verifies ID Token securely
✅ Fetches and displays user claims (email, name, etc.)
✅ Uses standard Go libraries (net/http) and coreos/go-oidc
-----------------------------------------------
🧱 Tech Stack
Go 1.20+
Okta Developer Account
OIDC via coreos/go-oidc
OAuth2 via golang.org/x/oauth2
-----------------------------------------------
🔧 Prerequisites
Okta Developer Account
Sign up at https://developer.okta.com
Create OIDC App in Okta:
Go to Applications → Create App Integration
Select OIDC - Web Application
Login redirect URI: http://localhost:8080/callback
Logout redirect URI: http://localhost:8080
Save and note:
Client ID
Client Secret
Okta domain (e.g., https://dev-123456.okta.com)
-----------------------------------------------
📦 Installation
git clone https://github.com/sharmabhishek15/oidc-okta-go-app.git
cd oidc-okta-go-app
go mod tidy
-----------------------------------------------
Note - Use deatils according to your OKTA account.
