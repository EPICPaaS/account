package filter

import (
	"fmt"
	"github.com/EPICPaaS/account/tools"
	"github.com/astaxie/beego/context"
)

func HandleAccess(ctx *context.Context) {
	token := ctx.GetCookie("epic_user_token")
	fmt.Println("获取到的token-" + token)
	ok, _ := tools.VerifyToken(token)
	if len(token) != 0 && ok {
		ctx.Redirect(302, "/succeed")
	}
}
