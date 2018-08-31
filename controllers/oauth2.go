package controllers

import (
	"context"
	"github.com/astaxie/beego"
	"github.com/kryazhev/blockchain-su/models"
	"golang.org/x/oauth2"
	"regexp"
	"strings"
)

var oauthConfigs = make(map[string]*oauth2.Config)

func init() {
	endpointNames := models.LookupEnv("oauth2-endpoint-names", "")

	if len(endpointNames) > 0 {
		split := strings.Split(endpointNames, ",")
		for _, endpointName := range split {
			config, err := models.OAuthConfig(endpointName)

			if err != nil {
				panic(err)
			}

			beego.Trace("Loading OAuth2.0 for ", endpointName, config)
			oauthConfigs[endpointName] = config
		}
	}
}

func (c *AppController) Login() {
	session := c.StartSession()

	user := models.NewUser(c.GetString("email", ""))
	session.Set("user", user)

	page := c.GetString("page")
	c.initDataWithUser(page, user)
}

func (c *AppController) Callback() {
	state := c.GetString("state")

	var user *models.User
	m := regexp.MustCompile(`^([a-z]+)`).FindStringSubmatch(state)
	if len(m) > 1 {
		endpointName := m[1]
		code := c.GetString("code")

		config := oauthConfigs[endpointName]

		// TODO setup timeout settings
		token, err := config.Exchange(context.TODO(), code)
		if err != nil {
			beego.Error(err)
		}

		session := c.StartSession()

		// TODO get userInfo (email, picture)
		user = models.NewUser(token.TokenType)
		session.Set("user", user)
	}

	page := c.GetString("page")
	c.initDataWithUser(page, user)
}

func (c *AppController) Logout() {
	beego.GlobalSessions.SessionDestroy(c.Ctx.ResponseWriter, c.Ctx.Request)

	c.initDataWithUser(c.GetString("page"), nil)
}
