package controllers

import (
	"github.com/astaxie/beego"
	"github.com/kryazhev/blockchain-su/models"
	"github.com/kryazhev/oauth2"
	"regexp"
)

var languages = []string{
	"ru", "az", "am", "by", "ge",
	"kz", "kg", "lv", "lt", "md",
	"tj", "tm", "ua", "uz", "ee",
	"en"}

/* OAuth 2.0 */
func (c *AppController) Callback() {
	state := c.GetString("state")

	var user *oauth2.User
	var err error
	m := regexp.MustCompile(`^([a-z]+)`).FindStringSubmatch(state)
	if len(m) > 1 {
		endpointName := m[1]
		code := c.GetString("code")

		config := oauth2.AuthConfigs[endpointName]

		user, err = config.GetUser(endpointName, code)

		if err != nil {
			beego.Error(err)
		} else {
			session := c.StartSession()
			session.Set("user", user)
		}
	}

	page := c.GetString("page")
	c.initDataWithUser(page, user)
}

func (c *AppController) Logout() {
	beego.GlobalSessions.SessionDestroy(c.Ctx.ResponseWriter, c.Ctx.Request)

	c.initDataWithUser(c.GetString("page"), nil)
}

/* Common */
func (c *AppController) ChangeLanguage() {
	page := c.GetString("page")
	c.initData(page)
}

func (c *AppController) Feedback() {
	err := models.SendEmail(c.GetString("email"), "vassiliy.kryazhev@gmail.com", "Feedback", c.GetString("message"))

	if err == nil {
		c.ajaxResponseSuccess(nil)
	} else {
		beego.Error("Can`t sent email ", err)
		c.ajaxResponseFail(err.Error())
	}
}

/* Developer Tools */
func (c *AppController) Login() {
	session := c.StartSession()

	email := c.GetString("email", "")
	user := &oauth2.User{Email: email}
	session.Set("user", user)

	page := c.GetString("page")
	c.initDataWithUser(page, user)
}

func (c *AppController) Header() {
	c.ajaxResponseSuccess(c.Ctx.Request.Header)
}

func (c *AppController) Sessions() {
	c.ajaxResponseSuccess(beego.GlobalSessions.GetActiveSession())
}
