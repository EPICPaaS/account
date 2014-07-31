package routers

import (
	"github.com/EPICPaaS/account/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{}, "get:Get;post:Login")
	beego.Router("/register", &controllers.RegisterController{}, "get:Get;post:Register")
	beego.Router("/succeed", &controllers.RegisterController{}, "get:Succeed")

}
