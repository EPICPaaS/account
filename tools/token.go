package tools

import (
	"encoding/base64"
	"github.com/astaxie/beego/utils"
	"net/url"
	"strconv"
	"time"
)

func CreateToken(userId string) string {
	values := make(url.Values, 3)
	values.Add("userid", userId)
	values.Add("secret", string(utils.RandomCreateBytes(20)))
	date := time.Time{}
	values.Add("time", strconv.FormatInt(int64(date.Nanosecond()), 10))
	token := base64.URLEncoding.EncodeToString([]byte(values.Encode()))
	return token
}
