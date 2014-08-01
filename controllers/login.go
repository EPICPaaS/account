package controllers

import (
	"github.com/EPICPaaS/account/modules/auth"
	"github.com/EPICPaaS/account/tools"
	"github.com/astaxie/beego"
	"strconv"
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
	ok, user := auth.VerifyUser(username, password)
	if !ok {
		this.TplNames = "login.html"
		this.Data["error"] = "用户名或密码错误!"
		this.Data["epic_sub_site"] = loginRedirect
		this.Data["UserName"] = username
		return
	}
	//生成用户登录token
	token, err := tools.CreateToken(strconv.Itoa(user.Id))
	if len(token) == 0 || err != nil {
		this.TplNames = "login.html"
		this.Data["error"] = "生成Token失败"
		this.Data["epic_sub_site"] = loginRedirect
		this.Data["UserName"] = username
		return
	}
	this.Ctx.SetCookie("epic_user_token ", token)
	this.Data["token"] = token
	this.Data["epic_sub_site"] = loginRedirect

	subSitesConf := beego.AppConfig.String("sub-sites")

	this.Data["srcs"] = strings.Split(subSitesConf, ",")

	this.TplNames = "loginRedirect.html"
}

func (this *LoginController) LoginOut() {
	token := this.GetString("token")
	if len(token) == 0 {
		token = this.Ctx.GetCookie("epic_user_token")
	}
	if len(token) != 0 {
		tools.DeleteToken(token)
	}
	this.Ctx.SetCookie("epic_user_token ", "")
	this.TplNames = "login.html"
}
