package main

import (
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
	_ "github.com/kryazhev/blockchain-su/controllers"
	"github.com/kryazhev/blockchain-su/models"
	_ "github.com/kryazhev/blockchain-su/routers"
	"os"
	"strings"
)

func main() {
	// Display list of environment variables
	for _, env := range []string{"PORT", "smtp.address", "smtp.host", "smtp.user", "smtp.password"} {
		beego.Trace(env + "=" + os.Getenv(env))
	}

	// Initialize language type list.
	langTypes := strings.Split(beego.AppConfig.String("lang::types"), "|")

	// Load locale files according to language types.
	for _, lang := range langTypes {
		beego.Trace("Loading language:", lang)
		if err := i18n.SetMessage(lang, "static/i18n/locale_"+lang+".ini"); err != nil {
			panic(err)
		}
	}

	//beego.AppConfig.Set("sessionConfig", `{"cookieName": "gosessionid", "cclifetime": 3600}`)

	beego.AddFuncMap("i18n", i18n.Tr)
	beego.AddFuncMap("safeHtml", models.SafeHtml)
	beego.AddFuncMap("data", models.Data)
	beego.AddFuncMap("lookupEnv", models.LookupEnv)
	beego.Run()
}
