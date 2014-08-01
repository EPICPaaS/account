package models

import (
	"github.com/EPICPaaS/account/tools"
	"github.com/astaxie/beego/orm"
	"time"
)

type User struct {
	Id          int
	UserName    string `orm:"size(30);unique"`
	NickName    string `orm:"size(30)"`
	Password    string `orm:"size(128)"`
	Url         string `orm:"size(100)"`
	Company     string `orm:"size(30)"`
	Location    string `orm:"size(30)"`
	Email       string `orm:"size(80);unique"`
	GrEmail     string `orm:"size(32)"`
	Info        string ``
	Github      string `orm:"size(30)"`
	Twitter     string `orm:"size(30)"`
	Google      string `orm:"size(30)"`
	Weibo       string `orm:"size(30)"`
	Linkedin    string `orm:"size(30)"`
	Facebook    string `orm:"size(30)"`
	PublicEmail bool   ``
	Followers   int    ``
	Following   int    ``
	FavTopics   int    ``
	IsAdmin     bool   `orm:"index"`
	IsActive    bool   `orm:"index"`
	IsForbid    bool   `orm:"index"`
	Lang        int    `orm:"index"`
	//	LangAdds    SliceStringField `orm:"size(50)"`
	Rands        string    `orm:"size(10)"`
	Created      time.Time `orm:"auto_now_add"`
	Updated      time.Time `orm:"auto_now"`
	QQ           string    `orm:"size(20)"`
	TelNum       string    `orm:"size(30)"`
	UserSocialId string    ``
	Identify     string    ``
}

func (m *User) Insert() error {
	m.Rands = GetUserSalt()
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *User) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *User) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *User) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *User) Exists() bool {
	num := 0
	orm.NewOrm().Raw("select 1 from user where user_name=? or email=?", m.UserName, m.Email).QueryRow(&num)
	if num == 1 {
		return true
	} else {
		return false
	}
}

// return a user salt token
func GetUserSalt() string {
	return tools.GetRandomString(10)
}
