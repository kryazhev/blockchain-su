package controllers

/* Common */
func (c *AppController) Get() {
	c.render("home")
}

func (c *AppController) AboutUs() {
	c.render("about-us")
}

func (c *AppController) ContactUs() {
	c.render("contact-us")
}

/* Projects */
func (c *AppController) HousingCooperative() {
	c.render("project.housing-cooperative")
}

func (c *AppController) Ussr() {
	c.render("project.ussr-2.0")
}

func (c *AppController) PensionFund() {
	c.render("project.pension-fund")
}

func (c *AppController) MunicipalServices() {
	c.render("project.municipal-services")
}

func (c *AppController) Bank() {
	c.render("project.bank")
}
