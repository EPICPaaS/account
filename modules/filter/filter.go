package filter

import (
	"github.com/EPICPaaS/account/tools"
	"github.com/astaxie/beego/context"
)

func HandleAccess(ctx *context.Context) {
	token := ctx.GetCookie("epic_user_token")
	ok, _ := tools.VerifyToken(token)
	if len(token) != 0 && ok {
		ctx.Redirect(302, "/succeed")
	}
}
