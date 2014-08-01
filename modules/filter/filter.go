package filter

import (
	"fmt"
	"github.com/astaxie/beego/context"
)

func HandleAccess(ctx *context.Context) {
	token := ctx.GetCookie("epic_user_token")

	fmt.Println("获取到的cookie为:" + token)
	if len(token) != 0 {
		ctx.Redirect(302, "/succeed")
	}
}
