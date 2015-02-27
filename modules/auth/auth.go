package auth

import (
	"encoding/json"
	"fmt"
	"github.com/EPICPaaS/account/models"
	"github.com/EPICPaaS/account/tools"
	"github.com/astaxie/beego/orm"
	"strconv"
	"strings"
)

func RegisterUser(user *models.User, username, email, password string) error {
	// use random salt encode password
	salt := models.GetUserSalt()
	pwd := tools.EncodePassword(password, salt)
	user.UserName = strings.ToLower(username)
	user.Email = strings.ToLower(email)
	// save salt and encode password, use $ as split char
	user.Password = fmt.Sprintf("%s$%s", salt, pwd)
	// save md5 email value for gravatar
	user.GrEmail = tools.EncodeMd5(user.Email)

	// Use username as default nickname.
	user.NickName = user.UserName
	//设置用户默认激活
	user.IsActive = true

	return user.Insert()
}

func ConnectUpdateUser(user *models.User, password string) error {
	salt := models.GetUserSalt()
	pwd := tools.EncodePassword(password, salt)
	user.Password = fmt.Sprintf("%s$%s", salt, pwd)
	return user.Update("UserName", "Password")
}

func UserIsExists(username, email string) bool {
	user := models.User{}
	user.UserName = strings.ToLower(username)
	user.Email = strings.ToLower(email)
	return user.Exists()
}

func InitConnect(identify string) (string, bool) {
	user := models.User{}
	user.Identify = identify
	err := user.Read("Identify")
	if err != nil {
		err = user.Insert()
		if err != nil {
			fmt.Println("connect创建用户失败-" + err.Error())
		}
	}
	id := user.Id
	password := user.Password
	if len(password) == 0 {
		return strconv.Itoa(id), false
	} else {
		return strconv.Itoa(id), true
	}

}

func VerifyUser(username, password string) (bool, *models.User) {
	isExists := UserIsExists(username, username)
	user := models.User{}
	if !isExists {
		return false, &user
	}
	var err error
	qs := orm.NewOrm()
	if strings.IndexRune(username, '@') == -1 {
		user.UserName = username
		err = qs.Read(&user, "UserName")
	} else {
		user.Email = username
		err = qs.Read(&user, "Email")
	}
	if err != nil {
		fmt.Println("用户登录读取用户信息失败" + err.Error())
		return false, &user
	}

	ok := VerifyPassword(password, user.Password)
	return ok, &user
}

func VerifyPassword(rawPwd, encodedPwd string) bool {
	var salt, encoded string
	if len(encodedPwd) > 11 {
		salt = encodedPwd[:10]
		encoded = encodedPwd[11:]
	}
	return tools.EncodePassword(rawPwd, salt) == encoded
}

func GetUserInfo(userid string) (bool, models.User) {
	var err error
	user := models.User{}
	user.Id, _ = strconv.Atoi(userid)
	userKey := "$account_userid_" + strconv.Itoa(user.Id)
	value, err := tools.RedisStorageInstance.GetKey(userKey)
	if len(value) != 0 || err == nil {
		valueByte := []byte(value)
		err := json.Unmarshal(valueByte, &user)
		if err == nil {
			return true, user
		}
	}
	qs := orm.NewOrm()
	err = qs.Read(&user, "Id")
	if err != nil {
		fmt.Println("用户登录读取用户信息失败" + err.Error())
		return false, user
	}
	userByte, _ := json.Marshal(&user)
	tools.RedisStorageInstance.SetKey(userKey, string(userByte))
	return true, user

}

func GetUserInfoFrmDB(userid string) (bool, models.User) {
	user := models.User{}
	user.Id, _ = strconv.Atoi(userid)
	qs := orm.NewOrm()
	err := qs.Read(&user, "Id")
	if err != nil {
		fmt.Println("用户登录读取用户信息失败" + err.Error())
		return false, user
	}
	return true, user
}

func SaveNewPassword(user *models.User, password string) error {
	salt := models.GetUserSalt()
	user.Password = fmt.Sprintf("%s$%s", salt, tools.EncodePassword(password, salt))
	return user.Update("Password", "Rands", "Updated")
}
