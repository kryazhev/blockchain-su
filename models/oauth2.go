package models

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/astaxie/beego"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	"golang.org/x/oauth2/github"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/odnoklassniki"
	"golang.org/x/oauth2/vk"
	"net/http"
	"os"
	"strings"
	"time"
)

type Config struct {
	DataURL string
	oauth2.Config
}

var httpClient = &http.Client{Timeout: 5 * time.Second}

var AuthConfigs = make(map[string]*Config)

func init() {
	endpointNames := LookupEnv("oauth2-endpoint-names", "")

	if len(endpointNames) > 0 {
		split := strings.Split(endpointNames, ",")
		for _, endpointName := range split {
			config, err := OAuthConfig(endpointName)

			if err != nil {
				panic(err)
			}

			beego.Trace("Loading OAuth2.0 for ", endpointName, config)
			AuthConfigs[endpointName] = config
		}
	}
}

func OAuthConfig(endpointName string) (*Config, error) {
	var endpoint oauth2.Endpoint

	switch endpointName {
	case "google":
		endpoint = google.Endpoint
	case "facebook":
		endpoint = facebook.Endpoint
	case "odnoklassniki":
		endpoint = odnoklassniki.Endpoint
	case "vk":
		endpoint = vk.Endpoint
	case "github":
		endpoint = github.Endpoint
	default:
		return nil, errors.New("Init unknown OAuth2.0 Endpoint : " + endpointName)
	}

	scopes := LookupEnv(endpointName+".oauth2.scopes", "openid email")

	return &Config{
		DataURL: os.Getenv(endpointName + ".oauth2.data-url"),
		Config: oauth2.Config{
			RedirectURL:  os.Getenv("oauth2-redirect-uri"),
			ClientID:     os.Getenv(endpointName + ".oauth2.client-id"),
			ClientSecret: os.Getenv(endpointName + ".oauth2.secret"),
			Endpoint:     endpoint,
			Scopes:       strings.Split(scopes, ",")}}, nil
}

func (cfg *Config) UserInfo(endpointName string, code string) (*User, error) {
	ctx := context.WithValue(context.Background(), oauth2.HTTPClient, httpClient)
	token, err := cfg.Exchange(ctx, code)

	if err != nil {
		return nil, err
	}

	beego.Trace(cfg.DataURL + token.AccessToken)

	client := cfg.Client(ctx, token)
	response, err := client.Get(cfg.DataURL + token.AccessToken)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	var data map[string]interface{}

	err = json.NewDecoder(response.Body).Decode(&data)

	if err != nil {
		return nil, err
	}

	beego.Trace(data)

	switch endpointName {
	case "google":
		return newGoogleUser(data), nil
	case "facebook":
		return newFacebookUser(data), nil
	case "github":
		return newGithubUser(data), nil
	default:
		return nil, errors.New("UserInfo unknown OAuth2.0 Endpoint : " + endpointName)
	}
}

func newGoogleUser(source map[string]interface{}) *User {
	return &User{
		Name:    source["name"].(string),
		Email:   source["email"].(string),
		Picture: source["picture"].(string)}
}

func newFacebookUser(source map[string]interface{}) *User {
	picture := source["picture"].(map[string]interface{})
	data := picture["data"].(map[string]interface{})

	return &User{
		Name:    source["name"].(string),
		Email:   source["email"].(string),
		Picture: data["url"].(string)}
}

func newGithubUser(source map[string]interface{}) *User {
	// TODO fix null values
	return &User{
		Name:    source["login"].(string),
		Picture: source["avatar_url"].(string)}
}
