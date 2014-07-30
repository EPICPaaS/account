package controllers

import (
	"fmt"
	"github.com/EPICPaaS/account/models"
	"github.com/EPICPaaS/account/modules/auth"
	"github.com/EPICPaaS/account/setting"
	"github.com/astaxie/beego"
)

type RegisterController struct {
	beego.Controller
}

func (this *RegisterController) Get() {
	this.Data["AppUrl"] = beego.AppConfig.String("appUrl")
	fmt.Println(setting.Captcha.VerifyReq(this.Ctx.Request))
	this.TplNames = "register.html"
}

func (this *RegisterController) Register() {
	user := models.User{}
	err := this.ParseForm(&user)
	this.TplNames = "succeed.html"
	if err != nil {
		beego.Error("注册失败-表单解析出错", err)
		this.Data["state"] = "注册失败"
		this.Data["msg"] = err.Error()
		return
	}
	isExist := auth.UserIsExists(user.UserName, user.Email)
	if isExist {
		this.Data["state"] = "注册失败"
		this.Data["msg"] = "[用户名]或者[邮箱]已被注册"
		return
	}
	err = auth.RegisterUser(&user, user.UserName, user.Email, user.Password)
	if err != nil {
		this.Data["state"] = "注册失败"
		beego.Error("注册失败-插入数据库出错", err)
		this.Data["msg"] = err.Error()
		return
	}
	this.Data["state"] = "注册成功"
	this.Data["msg"] = "恭喜"

}
func (this *RegisterController) Succeed() {
	this.Data["state"] = "注册成功"
	this.Data["msg"] = "用户已经存在"
	this.TplNames = "succeed.html"
}