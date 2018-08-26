package controllers

import (
	"github.com/astaxie/beego"
	"github.com/kryazhev/blockchain-su/models"
)

type AppController struct {
	beego.Controller
}

func (c *AppController) AjaxResponseSuccess(data interface{}) {
	c.Data["json"] = models.Result{Success: true, Data: data}
	c.ServeJSON()
}

func (c *AppController) AjaxResponseFail(message string) {
	c.Data["json"] = models.Result{Success: false, Message: message}
	c.ServeJSON()
}

func (c *AppController) Prepare() {
	if c.Data["Lang"] == nil {
		lang := c.Ctx.GetCookie("lang")
		if lang == "" {
			lang = "ru"
		}
		c.Data["Lang"] = lang
	}
}
