package routers

import (
	"github.com/EPICPaaS/account/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/account", &controllers.MainController{})
	beego.Router("/account/login", &controllers.LoginController{}, "get:Get;post:Login")
	beego.Router("/account/register", &controllers.RegisterController{}, "get:Get;post:Register")
	beego.Router("/account/succeed", &controllers.RegisterController{}, "get:Succeed")

}
