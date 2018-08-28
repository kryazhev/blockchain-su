package routers

import (
	"github.com/astaxie/beego"
	"github.com/kryazhev/blockchain-su/controllers"
)

func init() {
	beego.Router("/", &controllers.AppController{})
	beego.Router("/home.html", &controllers.AppController{}, "get:Get")
	beego.Router("/about-us.html", &controllers.AppController{}, "get:AboutUs")
	beego.Router("/contact-us.html", &controllers.AppController{}, "get:ContactUs")

	/* projects */
	beego.Router("/project/housing-cooperative.html", &controllers.AppController{}, "get:HousingCooperative")
	beego.Router("/project/ussr-2.0.html", &controllers.AppController{}, "get:Ussr")
	beego.Router("/project/pension-fund.html", &controllers.AppController{}, "get:PensionFund")
	beego.Router("/project/municipal-services.html", &controllers.AppController{}, "get:MunicipalServices")
	beego.Router("/project/bank.html", &controllers.AppController{}, "get:Bank")

	/* actions */
	beego.Router("/action/login", &controllers.AppController{}, "post:Login")
	beego.Router("/action/logout", &controllers.AppController{}, "get:Logout")

	beego.Router("/action/change-language", &controllers.AppController{}, "get:ChangeLanguage")
	beego.Router("/action/feedback", &controllers.AppController{}, "post:Feedback")
	beego.Router("/action/header", &controllers.AppController{}, "get:Header")
}
