package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	oidc "github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
)

var (
	clientID     = os.Getenv("OKTA_CLIENT_ID")
	clientSecret = os.Getenv("OKTA_CLIENT_SECRET")
	oktaDomain   = os.Getenv("OKTA_DOMAIN") // e.g., https://dev-xxxx.okta.com

	redirectURL = "http://localhost:8080/authorization-code/callback"
	provider    *oidc.Provider
	oauth2Config *oauth2.Config
	verifier     *oidc.IDTokenVerifier
)

func main() {
	ctx := context.Background()

	var err error
	provider, err = oidc.NewProvider(ctx, oktaDomain+"/oauth2/default")
	if err != nil {
		log.Fatalf("Failed to get provider: %v", err)
	}

	oidcConfig := &oidc.Config{
		ClientID: clientID,
	}

	verifier = provider.Verifier(oidcConfig)

	oauth2Config = &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Endpoint:     provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}

	http.HandleFunc("/", handleHome)
	http.HandleFunc("/login", handleLogin)
	http.HandleFunc("/callback", handleCallback)

	fmt.Println("Listening on http://localhost:8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, `<a href="/login">Login with Okta</a>`)
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	state := "random-state" // You should generate and validate this
	http.Redirect(w, r, oauth2Config.AuthCodeURL(state), http.StatusFound)
}

func handleCallback(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	code := r.URL.Query().Get("code")
	oauth2Token, err := oauth2Config.Exchange(ctx, code)
	if err != nil {
		http.Error(w, "Failed to exchange token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	rawIDToken, ok := oauth2Token.Extra("id_token").(string)
	if !ok {
		http.Error(w, "No id_token field in oauth2 token.", http.StatusInternalServerError)
		return
	}

	idToken, err := verifier.Verify(ctx, rawIDToken)
	if err != nil {
		http.Error(w, "Failed to verify ID Token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	var claims map[string]interface{}
	if err := idToken.Claims(&claims); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Login successful!\n\n")
	for k, v := range claims {
		fmt.Fprintf(w, "%s: %v\n", k, v)
	}
}
