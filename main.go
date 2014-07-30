package main

import (
	_ "github.com/EPICPaaS/account/routers"
	"github.com/EPICPaaS/account/setting"
	"github.com/astaxie/beego"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	beego.SetStaticPath("/account/static", "static")
	setting.LoadConfig()
	beego.Run()
}
