package tools

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/utils"
	"strconv"
	"time"
)

var tokenExpireTIme = 1800

type Token struct {
	UserId string `json:"userId"`
	Secret string `json:"secret"`
	Time   string `json:"time"`
}

func CreateToken(userId string) (string, error) {
	token := Token{}
	token.UserId = userId
	token.Secret = string(utils.RandomCreateBytes(20))
	token.Time = strconv.FormatInt(time.Now().UnixNano(), 10)
	tokenByte, _ := json.Marshal(&token)
	tokenStr := base64.URLEncoding.EncodeToString(tokenByte)
	_, err := RedisStorageInstance.SetExpireKey(tokenStr, tokenStr, tokenExpireTIme)
	if err != nil {
		fmt.Println("redis操作失败" + err.Error())
		return "", err
	}
	return tokenStr, nil
}

func VerifyToken(token string) (bool, string) {
	if len(token) == 0 {
		return false, ""
	}
	//1.判断token是否有效
	value, err := RedisStorageInstance.GetKey(token)
	if len(value) == 0 || err != nil {
		fmt.Println("token校验失败：" + err.Error())
		return false, ""
	}
	RedisStorageInstance.ExpireKey(token, tokenExpireTIme)
	//2.解密token获取userID
	tokenByte, _ := base64.URLEncoding.DecodeString(token)
	tokenObj := Token{}
	json.Unmarshal(tokenByte, &tokenObj)
	userId := tokenObj.UserId
	return true, userId
}

func DeleteToken(token string) bool {
	_, err := RedisStorageInstance.DelKey(token)
	if err != nil {
		fmt.Println("注销用户失败" + err.Error())
		return false
	} else {
		return true
	}
}
