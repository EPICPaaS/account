package main

import (
	_ "github.com/EPICPaaS/account/routers"
	"github.com/EPICPaaS/account/setting"
	"github.com/EPICPaaS/account/tools"
	"github.com/astaxie/beego"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	beego.SetStaticPath("/static", "static")
	setting.LoadConfig()
	tools.InitRedis()
	beego.SessionOn = true
	beego.Run()
}
