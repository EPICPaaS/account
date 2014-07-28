package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {

	this.Data["AppUrl"] = beego.AppConfig.String("appUrl")
	this.TplNames = "login.html"
}
