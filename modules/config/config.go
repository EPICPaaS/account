package config

import (
	"fmt"
	"github.com/EPICPaaS/account/models"
	"github.com/EPICPaaS/account/tools"
)

func GetSubSites() string {
	redisKey := "account-sub-sites"
	value, _ := tools.RedisStorageInstance.GetKey(redisKey)
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
	tools.RedisStorageInstance.SetKey(redisKey, config.Value)
	return config.Value
}
