package controllers

import (
	"github.com/EPICPaaS/account/modules/auth"
	"github.com/EPICPaaS/account/tools"
	"github.com/astaxie/beego"
)

type VerifyToken struct {
	beego.Controller
}

type VerifyTokenResult struct {
	Succeed  bool   `json:"succeed"`
	Userid   string `json:"userid"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func (this *VerifyToken) Get() {
	token := this.GetString("token")
	result := VerifyTokenResult{}
	if len(token) == 0 {
		result.Succeed = false
		this.Data["json"] = &result
		this.ServeJson()
		return
	}
	ok, userid := tools.VerifyToken(token)
	if !ok {
		result.Succeed = false
		this.Data["json"] = &result
		this.ServeJson()
		return
	}
	ok, user := auth.GetUserInfo(userid)
	if ok {
		result.Username = user.UserName
		result.Email = user.Email
	}
	result.Succeed = true
	result.Userid = userid
	this.Data["json"] = &result
	this.ServeJson()
	return
}

func (this *VerifyToken) Post() {
	this.Get()
}
