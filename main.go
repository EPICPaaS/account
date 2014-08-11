package main

import (
	_ "github.com/EPICPaaS/account/routers"
	"github.com/EPICPaaS/account/setting"
	"github.com/EPICPaaS/account/tools"
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/session/redis"
	_ "github.com/garyburd/redigo/redis"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	beego.SetStaticPath("/static", "static")
	tools.InitRedis()
	setting.LoadConfig()
	beego.SessionOn = true
	beego.SessionProvider = "redis"
	beego.SessionSavePath = beego.AppConfig.String("redis_resource")
	beego.Run()
}
