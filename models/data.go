package models

import (
	"errors"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	"golang.org/x/oauth2/github"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/odnoklassniki"
	"os"
	"strings"
)

type Result struct {
	Success bool        `json:"success,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type User struct {
	Email string
}

func NewUser(email string) *User {
	return &User{Email: email}
}

func (u *User) Name() string {
	data := strings.Split(u.Email, "@")
	if data != nil {
		return data[0]
	} else {
		return ""
	}
}

func OAuthConfig(endpointName string) (*oauth2.Config, error) {
	var endpoint oauth2.Endpoint

	switch endpointName {
	case "google":
		endpoint = google.Endpoint
	case "facebook":
		endpoint = facebook.Endpoint
	case "odnoklassniki":
		endpoint = odnoklassniki.Endpoint
	case "github":
		endpoint = github.Endpoint
	default:
		return nil, errors.New("Unknown OAuth2.0 Endpoint: " + endpointName)
	}

	scopes := LookupEnv(endpointName+".oauth2.scopes", "openid email")

	return &oauth2.Config{
		RedirectURL:  os.Getenv("oauth2-redirect-uri"),
		ClientID:     os.Getenv(endpointName + ".oauth2.client-id"),
		ClientSecret: os.Getenv(endpointName + ".oauth2.secret"),
		Endpoint:     endpoint,
		Scopes:       strings.Split(scopes, ","),
	}, nil
}
