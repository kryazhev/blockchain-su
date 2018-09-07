package controllers

import (
	"github.com/astaxie/beego"
	"github.com/kryazhev/blockchain-su/models"
	"github.com/kryazhev/oauth2"
	"strings"
)

var pages = []string{
	"home", "about-us", "contact-us",
	"project.housing-cooperative", "project.ussr-2.0", "project.pension-fund", "project.municipal-services", "project.bank"}

type AppController struct {
	beego.Controller
}

func (c *AppController) initData(page string) {
	if !models.HasElem(pages, page) {
		page = "home"
	}

	c.Data["Page"] = page

	c.Data["PageTitle"] = "meta." + page + ".title"
	c.Data["PageDescription"] = "meta." + page + ".description"
	c.Data["PageKeywords"] = "meta." + page + ".keywords"

	c.TplName = strings.Replace(page, ".", "/", 1) + ".html"
}

func (c *AppController) initDataWithUser(page string, user *oauth2.User) {
	if user != nil {
		c.Data["User"] = user
	} else {
		delete(c.Data, "User")
	}

	c.initData(page)
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
	c.Data["OAuthConfigs"] = oauth2.AuthConfigs

	user := c.GetSession("user")
	if user != nil {
		c.Data["User"] = user
	}

	lang := c.GetString("lang", "")
	if len(lang) == 0 {
		lang = c.Ctx.GetCookie("lang")
	} else {
		c.Ctx.SetCookie("lang", lang)
	}
	if !models.HasElem(languages, lang) {
		lang = "ru"
	}
	c.Data["Lang"] = lang

	if lang == "en" {
		c.Data["Country"] = "us"
	} else {
		c.Data["Country"] = lang
	}

}
