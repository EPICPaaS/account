package controllers

import (
	"fmt"
	"github.com/EPICPaaS/account/models"
	"github.com/EPICPaaS/account/modules/auth"
	"github.com/EPICPaaS/account/tools"
	"github.com/astaxie/beego"
	"strconv"
)

type SocialAuthController struct {
	beego.Controller
}

func (this *SocialAuthController) ConnectPost() {
	userId := this.GetString("userId")
	password := this.GetString("Password")
	userName := this.GetString("UserName")
	if len(userId) == 0 || len(password) == 0 || len(userName) == 0 {
		this.Data["userId"] = userId
		this.TplNames = "connect.html"
		this.Data["error"] = "[用户名]或者[密码]为空"
		this.Data["state"] = "注册失败"
		this.Data["msg"] = "[用户名]或者[邮箱]已被注册"
		return
	}
	isExist := auth.UserIsExists(userName, userName)
	if isExist {
		this.Data["userId"] = userId
		this.TplNames = "connect.html"
		this.Data["error"] = "[用户名]或者[邮箱]已被注册"
		this.Data["state"] = "注册失败"
		this.Data["msg"] = "[用户名]或者[邮箱]已被注册"
		return
	}
	user := models.User{}
	user.Password = password
	user.UserName = userName
	user.Id, _ = strconv.Atoi(userId)
	err := auth.ConnectUpdateUser(&user, password)
	if err != nil {
		this.Data["userId"] = userId
		this.TplNames = "connect.html"
		this.Data["error"] = err.Error()
		this.Data["state"] = "注册失败"
		beego.Error("注册失败-插入数据库出错", err)
		this.Data["msg"] = err.Error()
		return
	}
	this.Data["state"] = "注册成功"
	this.Data["msg"] = "恭喜"
	this.TplNames = "succeed.html"
}

func (this *SocialAuthController) Connect() {
	identify := this.GetSession("custom_userSocial_identify")
	userId, isInit := auth.InitConnect(identify.(string))
	token, err := tools.CreateToken(userId)
	this.Data["userId"] = userId
	if err == nil {
		this.Ctx.SetCookie("epic_user_token", token)
	} else {
		fmt.Println("生成token失败-" + err.Error())
	}
	if isInit {
		this.Redirect("/login", 302)
	} else {
		this.TplNames = "connect.html"
	}

}
