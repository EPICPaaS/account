package tools

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/utils"
	"github.com/garyburd/redigo/redis"
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
	date := time.Time{}
	token.Time = strconv.FormatInt(int64(date.Nanosecond()), 10)
	tokenByte, _ := json.Marshal(&token)
	tokenStr := base64.URLEncoding.EncodeToString(tokenByte)
	fmt.Println("生成的token为-" + tokenStr)
	redisConn := RedisStorageInstance.getConn(tokenStr)
	defer redisConn.Close()
	err := redisConn.Send("SET", tokenStr, tokenStr)
	if err != nil {
		fmt.Println("redis操作失败" + err.Error())
		return "", err
	}
	err = redisConn.Send("EXPIRE", tokenStr, tokenExpireTIme)
	if err != nil {
		fmt.Println("redis操作失败" + err.Error())
		return "", err
	}
	err = redisConn.Flush()
	if err != nil {
		fmt.Println("redis操作失败" + err.Error())
		return "", err
	}
	fmt.Println("生成的token-" + tokenStr)
	return tokenStr, nil
}

func VerifyToken(token string) (bool, string) {
	//1.判断token是否有效
	redisConn := RedisStorageInstance.getConn(token)
	defer redisConn.Close()
	value, err := redis.String(redisConn.Do("GET", token))
	if len(value) == 0 || err != nil {
		return false, ""
	}
	//2.解密token获取userID
	tokenByte, _ := base64.URLEncoding.DecodeString(token)
	tokenObj := Token{}
	json.Unmarshal(tokenByte, &tokenObj)
	userId := tokenObj.UserId
	return true, userId
}

func DeleteToken(token string) bool {
	redisConn := RedisStorageInstance.getConn(token)
	defer redisConn.Close()
	_, err := redisConn.Do("DEL", token)
	if err != nil {
		fmt.Println("退出失败" + err.Error())
		return false
	} else {
		return true
	}
}
