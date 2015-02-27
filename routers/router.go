package routers

import (
	"github.com/EPICPaaS/account/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.LoginController{}, "get:Get;post:Login")
	beego.Router("/login", &controllers.LoginController{}, "get:Get;post:Login")
	beego.Router("/logout", &controllers.LoginController{}, "get:LoginOut;post:LoginOut")
	beego.Router("/verify_token", &controllers.VerifyToken{})
	beego.Router("/register", &controllers.RegisterController{}, "get:Get;post:Register")
	beego.Router("/succeed", &controllers.RegisterController{}, "get:Succeed")
	beego.Router("/register/connect", &controllers.SocialAuthController{}, "get:Connect;post:ConnectPost")
	beego.Router("/change/password", &controllers.SettingController{}, "get:ChangePassword;post:ChangePasswordSave")
}
