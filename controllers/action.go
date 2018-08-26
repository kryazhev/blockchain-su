package controllers

import (
	"github.com/astaxie/beego"
	"github.com/kryazhev/blockchain-su/models"
)

func (c *AppController) ChangeLanguage() {
	lang := c.GetString("lang", "ru")
	c.Ctx.SetCookie("lang", lang)

	c.Data["Lang"] = lang
	c.Data["Page"] = "home"
	c.TplName = "home.html"
}

func (c *AppController) Feedback() {
	err := models.SendEmail(c.GetString("email"), "vassiliy.kryazhev@gmail.com", "Feedback", c.GetString("message"))

	if err == nil {
		c.AjaxResponseSuccess(nil)
	} else {
		beego.Error("Can`t sent email ", err)
		c.AjaxResponseFail(err.Error())
	}
}

func (c *AppController) Request() {
	c.Data["json"] = c.Ctx.Request.Header
	c.ServeJSON()
}
