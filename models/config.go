package models

import (
	"github.com/astaxie/beego/orm"
)

type Config struct {
	Id    int
	Key   string `orm:"size(30)"`
	Value string `orm:"size(2048)"`
	Des   string `orm:"size(100)"`
}

func (m *Config) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}
