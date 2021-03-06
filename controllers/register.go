package controllers

import (
	"github.com/EPICPaaS/account/models"
	"github.com/EPICPaaS/account/modules/auth"
	"github.com/EPICPaaS/account/modules/config"
	"github.com/EPICPaaS/account/setting"
	"github.com/astaxie/beego"
)

type RegisterController struct {
	beego.Controller
}

func (this *RegisterController) Get() {
	this.Data["AppUrl"] = beego.AppConfig.String("appUrl")

	this.TplNames = "register.html"

	redirectURL := this.GetString("redirectURL")
	if "" == redirectURL {
		redirectURL = this.GetString("epic_sub_site")

		if "" == redirectURL {
			redirectURL = config.GetRedirectURL()
		}
	}

	this.Data["redirectURL"] = redirectURL
}

func (this *RegisterController) Register() {
	redirectURL := this.GetString("redirectURL")
	if "" == redirectURL {
		redirectURL = config.GetRedirectURL()
	}

	this.Data["redirectURL"] = redirectURL

	user := models.User{}
	err := this.ParseForm(&user)
	this.TplNames = "succeed.html"
	this.Data["succ"] = true
	if err != nil {
		beego.Error("注册失败-表单解析出错", err)
		this.Data["state"] = "注册失败"
		this.Data["msg"] = err.Error()
		this.Data["succ"] = false
		return
	}
	ok := setting.Captcha.VerifyReq(this.Ctx.Request)
	if !ok {
		this.Data["state"] = "注册失败"
		this.Data["msg"] = "验证码错误"
		this.Data["succ"] = false
		return
	}

	isExist := auth.UserIsExists(user.UserName, user.Email)
	if isExist {
		this.Data["state"] = "注册失败"
		this.Data["msg"] = "[用户名]或者[邮箱]已被注册"
		this.Data["succ"] = false
		return
	}
	err = auth.RegisterUser(&user, user.UserName, user.Email, user.Password)
	if err != nil {
		this.Data["state"] = "注册失败"
		beego.Error("注册失败-插入数据库出错", err)
		this.Data["msg"] = err.Error()
		this.Data["succ"] = false
		return
	}
	this.Data["state"] = "注册成功"
	this.Data["msg"] = "恭喜!!将进行自动跳转,请稍等..."

}
func (this *RegisterController) Succeed() {
	this.Data["appname"] = "单点登录服务"
	this.Data["state"] = "注册成功"
	this.Data["msg"] = "用户已经登录"
	this.TplNames = "succeed.html"
	this.Data["succ"] = true

	redirectURL := this.GetString("redirectURL")
	if "" == redirectURL {
		redirectURL = config.GetRedirectURL()
	}

	this.Data["redirectURL"] = redirectURL
}
