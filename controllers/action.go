package controllers

import (
	"github.com/astaxie/beego"
	"github.com/kryazhev/blockchain-su/models"
	"regexp"
)

var languages = []string{
	"ru", "az", "am", "by", "ge",
	"kz", "kg", "lv", "lt", "md",
	"tj", "tm", "ua", "uz", "ee",
	"us"}

/* OAuth 2.0 */
func (c *AppController) Callback() {
	state := c.GetString("state")

	var user *models.User
	var err error
	m := regexp.MustCompile(`^([a-z]+)`).FindStringSubmatch(state)
	if len(m) > 1 {
		endpointName := m[1]
		code := c.GetString("code")

		config := models.AuthConfigs[endpointName]

		user, err = config.UserInfo(endpointName, code)

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
	lang := c.GetString("lang", "ru")
	if !models.HasElem(languages, lang) {
		lang = "ru"
	}
	c.Ctx.SetCookie("lang", lang)
	c.Data["Lang"] = lang

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
	user := &models.User{Email: email}
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
