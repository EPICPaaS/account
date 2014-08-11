package config

import (
	"fmt"
	"github.com/EPICPaaS/account/models"
	"github.com/EPICPaaS/account/tools"
)

var (
	SubSitesRedisKey = "account-sub-sites"
)

func InitConfig() {
	config := models.Config{}
	config.Key = "sub-sites"
	err := config.Read("key")
	if err != nil {
		fmt.Println("获取配置【sub-sites】失败")
		return
	}
	tools.RedisStorageInstance.SetKey(SubSitesRedisKey, config.Value)
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
