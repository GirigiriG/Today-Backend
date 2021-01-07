package auth

import (
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	GoogleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:3001/secret",
		ClientID:     os.Getenv("AUTH_CLIENTID"),
		ClientSecret: os.Getenv("AUTH_CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}

	RandomState = os.Getenv("STATE_RANDOM")
)

//GoogleAuthResponse response from google auth
type GoogleAuthResponse struct {
	Id      string `json:id`
	Email   string `json:email`
	Picture string `json:picture`
	Token   string
}
