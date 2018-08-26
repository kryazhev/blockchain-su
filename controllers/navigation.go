package controllers

import (
	"github.com/astaxie/beego"
)

type AppController struct {
	beego.Controller
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

func (c *AppController) Get() {
	c.Data["Page"] = "home"
	c.TplName = "home.html"
}

func (c *AppController) AboutUs() {
	c.Data["Page"] = "about-us"
	c.TplName = "about-us.html"
}

func (c *AppController) ContactUs() {
	c.Data["Page"] = "contact-us"
	c.TplName = "contact-us.html"
}

/* Projects */
func (c *AppController) HousingCooperative() {
	c.Data["Page"] = "project.housing-cooperative"
	c.TplName = "project/housing-cooperative.html"
}

func (c *AppController) Ussr() {
	c.Data["Page"] = "project.ussr-2.0"
	c.TplName = "project/ussr-2.0.html"
}

func (c *AppController) PensionFund() {
	c.Data["Page"] = "project.pension-fund"
	c.TplName = "project/pension-fund.html"
}

func (c *AppController) MunicipalServices() {
	c.Data["Page"] = "project.municipal-services"
	c.TplName = "project/municipal-services.html"
}

func (c *AppController) Bank() {
	c.Data["Page"] = "project.bank"
	c.TplName = "project/bank.html"
}
