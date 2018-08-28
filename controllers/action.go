package controllers

import (
	"github.com/astaxie/beego"
	"github.com/kryazhev/blockchain-su/models"
)

var languages = []string{
	"ru", "az", "am", "by", "ge",
	"kz", "kg", "lv", "lt", "md",
	"tj", "tm", "ua", "uz", "ee",
	"us"}

func (c *AppController) Login() {
	session := c.StartSession()

	user := models.NewUser(c.GetString("email", ""))
	session.Set("user", user)

	c.initDataWithUser(c.GetString("page", "home"), user)
}

func (c *AppController) Logout() {
	session := c.StartSession()
	session.Delete("user")

	c.initDataWithUser(c.GetString("page", "home"), nil)
}

func (c *AppController) ChangeLanguage() {
	lang := c.GetString("lang", "ru")
	if !models.HasElem(languages, lang) {
		lang = "ru"
	}
	c.Ctx.SetCookie("lang", lang)
	c.Data["Lang"] = lang

	page := c.GetString("page", "home")
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

func (c *AppController) Request() {
	c.ajaxResponseSuccess(c.Ctx.Request.Header)
}
