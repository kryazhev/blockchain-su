package controllers

import (
	"github.com/astaxie/beego"
	"github.com/kryazhev/blockchain-su/models"
	"strings"
)

var pages = []string{
	"home", "about-us", "contact-us",
	"project.housing-cooperative", "project.ussr-2.0", "project.pension-fund", "project.municipal-services", "project.bank"}

type AppController struct {
	beego.Controller
}

func (c *AppController) render(page string) {
	c.Data["Page"] = page
	c.TplName = strings.Replace(page, ".", "/", 1) + ".html"
}

func (c *AppController) ajaxResponseSuccess(data interface{}) {
	c.Data["json"] = models.Result{Success: true, Data: data}
	c.ServeJSON()
}

func (c *AppController) ajaxResponseFail(message string) {
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
