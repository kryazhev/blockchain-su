package controllers

/* Base */
func (c *AppController) Get() {
	c.initData("home")
}

func (c *AppController) AboutUs() {
	c.initData("about-us")
}

func (c *AppController) ContactUs() {
	c.initData("contact-us")
}

/* Projects */
func (c *AppController) HousingCooperative() {
	c.initData("project.housing-cooperative")
}

func (c *AppController) Ussr() {
	c.initData("project.ussr-2.0")
}

func (c *AppController) PensionFund() {
	c.initData("project.pension-fund")
}

func (c *AppController) MunicipalServices() {
	c.initData("project.municipal-services")
}

func (c *AppController) Bank() {
	c.initData("project.bank")
}
