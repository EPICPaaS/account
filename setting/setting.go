package setting

import (
	"github.com/EPICPaaS/account/models"
	"github.com/EPICPaaS/account/modules/filter"
	"github.com/EPICPaaS/account/modules/socialAuth"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/utils/captcha"
	"github.com/beego/social-auth"
	"github.com/beego/social-auth/apps"
	_ "github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var (
	Captcha *captcha.Captcha
	Cache   cache.Cache
)

func init() {
	orm.RegisterModel(new(models.User), new(models.Config))
}

func LoadConfig() {
	redisServer := beego.AppConfig.String("redis_resource")
	redisServer = "{\"conn\":\"" + redisServer + "\"}"
	store, err := cache.NewCache("redis", redisServer)
	if err != nil {
		log.Panic("缓存出错" + err.Error())
	}
	Captcha = captcha.NewWithFilter("/captcha/", store)

	driverName := beego.AppConfig.String("driverName")
	dataSource := beego.AppConfig.String("dataSource")
	maxIdle, _ := beego.AppConfig.Int("maxIdle")
	maxOpen, _ := beego.AppConfig.Int("maxOpen")

	orm.RegisterDriver("mysql", orm.DR_MySQL)

	// set default database
	err = orm.RegisterDataBase("default", driverName, dataSource, maxIdle, maxOpen)
	if err != nil {
		beego.Error(err)
	}
	orm.RunCommand()

	err = orm.RunSyncdb("default", false, false)
	if err != nil {
		beego.Error(err)
	}
	SocialAuthInit()
}

func SocialAuthInit() {
	var clientId, secret string
	var err error
	appURL := beego.AppConfig.String("social_auth_url")
	if len(appURL) > 0 {
		social.DefaultAppUrl = appURL
	}

	clientId = beego.AppConfig.String("github_client_id")
	secret = beego.AppConfig.String("github_client_secret")
	err = social.RegisterProvider(apps.NewGithub(clientId, secret))
	if err != nil {
		beego.Error(err)
	}

	clientId = beego.AppConfig.String("google_client_id")
	secret = beego.AppConfig.String("google_client_secret")
	err = social.RegisterProvider(apps.NewGoogle(clientId, secret))
	if err != nil {
		beego.Error(err)
	}

	clientId = beego.AppConfig.String("weibo_client_id")
	secret = beego.AppConfig.String("weibo_client_secret")
	err = social.RegisterProvider(apps.NewWeibo(clientId, secret))
	if err != nil {
		beego.Error(err)
	}

	clientId = beego.AppConfig.String("qq_client_id")
	secret = beego.AppConfig.String("qq_client_secret")
	err = social.RegisterProvider(apps.NewQQ(clientId, secret))
	if err != nil {
		beego.Error(err)
	}

	socialAuth.SocialAuth = social.NewSocial("/register/connect", new(socialAuth.SocialAuther))
	socialAuth.SocialAuth.ConnectSuccessURL = "/register/connect"
	beego.InsertFilter("/login", beego.BeforeRouter, filter.HandleAccess)
	beego.InsertFilter("/", beego.BeforeRouter, filter.HandleAccess)
	beego.InsertFilter("/login/*/access", beego.BeforeRouter, socialAuth.HandleAccess)
	beego.InsertFilter("/login/*", beego.BeforeRouter, socialAuth.HandleRedirect)
}
