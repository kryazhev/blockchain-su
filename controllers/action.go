package controllers

import (
	"github.com/astaxie/beego"
	"github.com/kryazhev/blockchain-su/models"
)

/* Actions */
func (c *AppController) ChangeLanguage() {
	lang := c.GetString("lang", "ru")
	c.Ctx.SetCookie("lang", lang)

	c.Data["Lang"] = lang
	c.Data["Page"] = "home"
	c.TplName = "home.html"
}

func (c *AppController) Feedback() {
	from := c.GetString("email")
	message := c.GetString("message")

	var err error

	if message != "" {
		err = models.SendEmail(from, "vassiliy.kryazhev@gmail.com", "Feedback", message)
	}

	if err == nil {
		c.Data["json"] = models.Result{Success: true}
	} else {
		beego.Error("Can`t sent email ", err)
		c.Data["json"] = models.Result{Success: false, Message: err.Error()}
	}

	c.ServeJSON()
}

func (c *AppController) Request() {
	c.Data["json"] = c.Ctx.Request.Header
	c.ServeJSON()
}
