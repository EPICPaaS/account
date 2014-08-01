package controllers

import (
	"github.com/EPICPaaS/account/models"
	"github.com/EPICPaaS/account/modules/auth"
	"github.com/EPICPaaS/account/modules/socialAuth"
	"github.com/astaxie/beego"
)

type SocialAuthController struct {
	beego.Controller
}

func (this *SocialAuthController) ConnectPost() {
	userSocialId := this.GetString("id")
	identify := this.GetString("identify")
	password := this.GetString("Password")
	userName := this.GetString("UserName")
	this.TplNames = "succeed.html"
	isExist := auth.UserIsExists(userName, userName)
	if isExist {
		this.Data["state"] = "注册失败"
		this.Data["msg"] = "[用户名]或者[邮箱]已被注册"
		return
	}
	user := models.User{}
	user.UserSocialId = userSocialId
	user.Identify = identify
	err := auth.RegisterUser(&user, userName, "", password)
	if err != nil {
		this.Data["state"] = "注册失败"
		beego.Error("注册失败-插入数据库出错", err)
		this.Data["msg"] = err.Error()
		return
	}
	this.Data["state"] = "注册成功"
	this.Data["msg"] = "恭喜"
}

func (this *SocialAuthController) Connect() {

	st, ok := socialAuth.SocialAuth.ReadyConnect(this.Ctx)
	if !ok {
		this.Redirect("/login", 302)
		return
	}

	loginRedirect, userSocial, err := socialAuth.SocialAuth.ConnectAndLogin(this.Ctx, st, 1)
	if err != nil {
		// may be has error
		beego.Error(err)
		this.Redirect(loginRedirect, 302)
	} else {
		this.Data["identify"] = userSocial.Identify
		this.Data["id"] = userSocial.Id
	}
	this.TplNames = "connect.html"
}
