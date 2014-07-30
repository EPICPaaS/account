package auth

import (
	"fmt"
	"github.com/EPICPaaS/account/models"
	"github.com/EPICPaaS/account/tools"
	"github.com/astaxie/beego/orm"
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

func UserIsExists(username, email string) bool {
	user := models.User{}
	user.UserName = strings.ToLower(username)
	user.Email = strings.ToLower(email)
	return user.Exists()
}

func VerifyUser(username, password string) bool {
	isExists := UserIsExists(username, username)
	if !isExists {
		return false
	}
	user := models.User{}
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
		return false
	}

	return VerifyPassword(password, user.Password)
}

func VerifyPassword(rawPwd, encodedPwd string) bool {
	var salt, encoded string
	if len(encodedPwd) > 11 {
		salt = encodedPwd[:10]
		encoded = encodedPwd[11:]
	}
	return tools.EncodePassword(rawPwd, salt) == encoded
}
