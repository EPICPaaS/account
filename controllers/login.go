package controllers

import (
	"github.com/EPICPaaS/account/modules/auth"
	"github.com/EPICPaaS/account/tools"
	"github.com/astaxie/beego"
	"strings"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	this.Data["AppUrl"] = beego.AppConfig.String("appUrl")
	this.TplNames = "login.html"
	loginRedirect := strings.TrimSpace(this.GetString("epic_sub_site"))
	if tools.IsMatchHost(loginRedirect) == false {
		loginRedirect = "/"
	}
	if len(loginRedirect) > 0 {
		this.Data["epic_sub_site"] = loginRedirect
	}
}

func (this *LoginController) Login() {
	this.Data["AppUrl"] = beego.AppConfig.String("appUrl")
	username := this.GetString("UserName")
	password := this.GetString("Password")
	loginRedirect := this.GetString("epic_sub_site")
	ok := auth.VerifyUser(username, password)
	if !ok {
		this.TplNames = "login.html"
		this.Data["error"] = "用户名或密码错误!"
		this.Data["epic_sub_site"] = loginRedirect
		this.Data["UserName"] = username
		return
	}
	//生成用户登录token
	token := "asdfa"

	this.Data["token"] = token
	this.Data["epic_sub_site"] = loginRedirect
	this.TplNames = "loginRedirect.html"

}
