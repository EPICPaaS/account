package config

import (
	"fmt"

	"github.com/EPICPaaS/account/models"
	"github.com/EPICPaaS/account/tools"
)

var (
	SubSitesRedisKey = "account-sub-sites"
	RedirectURLKey   = "redirect-url"
)

func InitConfig() {
	// sub-sites
	config := models.Config{}
	config.Key = "sub-sites"
	err := config.Read("key")
	if err != nil {
		fmt.Println("获取配置【sub-sites】失败")
		return
	}
	tools.RedisStorageInstance.SetKey(SubSitesRedisKey, config.Value)

	// redirectURL
	config = models.Config{}
	config.Key = "redirectURL"
	err = config.Read("key")
	if err != nil {
		fmt.Println("获取配置【redirectURL】失败")
		return
	}

	tools.RedisStorageInstance.SetKey(RedirectURLKey, config.Value)
}

func GetSubSites() string {
	value, _ := tools.RedisStorageInstance.GetKey(SubSitesRedisKey)
	if len(value) != 0 {
		return value
	}
	config := models.Config{}
	config.Key = "sub-sites"
	err := config.Read("key")
	if err != nil {
		fmt.Println("获取配置【sub-sites】失败")
		return ""
	}
	tools.RedisStorageInstance.SetKey(SubSitesRedisKey, config.Value)
	return config.Value
}

func GetRedirectURL() string {
	value, _ := tools.RedisStorageInstance.GetKey(RedirectURLKey)
	if len(value) != 0 {
		return value
	}

	config := models.Config{}
	config.Key = "redirectURL"
	err := config.Read("key")
	if err != nil {
		fmt.Println("获取配置【redirectURL】失败")
		return ""
	}

	tools.RedisStorageInstance.SetKey(RedirectURLKey, config.Value)

	return config.Value
}
