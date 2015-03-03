package controllers

import (
	"github.com/EPICPaaS/account/modules/auth"
	"github.com/EPICPaaS/account/modules/config"
	"github.com/EPICPaaS/account/tools"
	"github.com/astaxie/beego"
)

type SettingController struct {
	beego.Controller
}

func (this *SettingController) ChangePassword() {

	this.Data["AppUrl"] = beego.AppConfig.String("appUrl")
	this.TplNames = "change_password.html"
	redirectURL := this.GetString("redirectURL")
	if "" == redirectURL {
		redirectURL = this.GetString("epic_sub_site")
		if "" == redirectURL {
			redirectURL = config.GetRedirectURL()
		}
	}
	this.Data["redirectURL"] = redirectURL
	this.Data["epic_sub_site"] = redirectURL
	ctx := this.Ctx
	token := ctx.GetCookie("epic_user_token")
	ok, _ := tools.VerifyToken(token)
	if len(token) == 0 || !ok {
		ctx.Redirect(302, "/")
		return
	}

}

func (this *SettingController) ChangePasswordSave() {
	redirectURL := this.GetString("epic_sub_site")
	if "" == redirectURL {
		redirectURL = config.GetRedirectURL()
	}
	this.Data["epic_sub_site"] = redirectURL
	this.Data["redirectURL"] = redirectURL
	this.Data["succ"] = false
	passwordOld := this.GetString("PasswordOld")
	password := this.GetString("Password")
	passwordRe := this.GetString("PasswordRe")
	if len(passwordOld) == 0 || len(password) == 0 || len(passwordRe) == 0 {
		this.Data["msg"] = "修改密码失败，缺少参数"
		this.TplNames = "change_password_succeed.html"
		return
	}
	if password != passwordRe {
		this.Data["msg"] = "修改密码失败，两次密码输入不一致"
		this.TplNames = "change_password_succeed.html"
		return
	}
	token := this.Ctx.GetCookie("epic_user_token")
	ok, userId := tools.VerifyToken(token)
	if len(token) == 0 || !ok {
		this.Data["msg"] = "修改密码失败，请重新登录"
		this.TplNames = "change_password_succeed.html"
		return
	}

	ok, user := auth.GetUserInfoFrmDB(userId)
	if !ok {
		this.Data["msg"] = "修改密码失败，用户不存在"
		this.TplNames = "change_password_succeed.html"
		return
	}
	ok = auth.VerifyPassword(passwordOld, user.Password)
	if !ok {
		this.Data["msg"] = "修改密码失败，当前密码验证错误"
		this.TplNames = "change_password_succeed.html"
		return
	}
	err := auth.SaveNewPassword(&user, password)
	if err != nil {
		beego.Error("密码修改失败:", err)
		this.Data["msg"] = "修改密码失败，请联系管理员"
		this.TplNames = "change_password_succeed.html"
		return
	}
	this.Data["msg"] = "修改密码成功，稍后将进行自动跳转"
	this.Data["succ"] = true
	this.TplNames = "change_password_succeed.html"
}
