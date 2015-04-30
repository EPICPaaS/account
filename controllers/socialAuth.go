package controllers

import (
	"fmt"
	"github.com/EPICPaaS/account/models"
	"github.com/EPICPaaS/account/modules/auth"
	"github.com/EPICPaaS/account/modules/config"
	"github.com/EPICPaaS/account/tools"
	"github.com/astaxie/beego"
	"strconv"
	"strings"
)

type SocialAuthController struct {
	beego.Controller
}

func (this *SocialAuthController) ConnectPost() {
	token := this.Ctx.GetCookie("epic_user_token")
	ok, userId := tools.VerifyToken(token)
	if !ok || len(userId) == 0 {
		this.Redirect("/", 302)
		return
	}
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

	subSitesConf := config.GetSubSites()
	this.Data["srcs"] = strings.Split(subSitesConf, ",")
	this.Data["token"] = token
	this.Data["state"] = "注册成功"
	this.Data["msg"] = "3秒后自动跳转!!"
	this.Data["succ"] = true
	this.Data["redirectURL"] = config.GetRedirectURL()
	this.TplNames = "succeed.html"
}

func (this *SocialAuthController) Connect() {
	identify := this.GetSession("custom_userSocial_identify")
	userName, ok := this.GetSession("custom_userSocial_userName").(string)

	if ok && len(userName) > 0 {
		userNames := strings.Split(userName, "_")
		this.Data["intrant"] = userNames[0] + "用户“" + userNames[1] + "”，"
	}
	userId, isInit := auth.InitConnect(identify.(string))
	token, err := tools.CreateToken(userId)
	this.Data["userId"] = userId
	if err == nil {
		this.Ctx.SetCookie("epic_user_token", token)
	} else {
		fmt.Println("生成token失败-" + err.Error())
	}
	if isInit {
		this.Data["token"] = token
		this.Data["epic_sub_site"] = config.GetRedirectURL()
		subSitesConf := config.GetSubSites()
		this.Data["srcs"] = strings.Split(subSitesConf, ",")
		this.TplNames = "loginRedirect.html"
	} else {
		this.TplNames = "connect.html"
	}

}
