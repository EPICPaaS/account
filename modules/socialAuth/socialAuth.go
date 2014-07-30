package socialAuth

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/beego/social-auth"
)

var SocialAuth *social.SocialAuth

type SocialAuther struct {
}

func (p *SocialAuther) IsUserLogin(ctx *context.Context) (int, bool) {
	ctx.GetCookie("epic_user_token")
	return 1, false
}

func (p *SocialAuther) LoginUser(ctx *context.Context, uid int) (string, error) {
	// fake login the user
	if uid == 1 {
		ctx.Input.CruSession.Set("login_user", 1)
	}
	return "/login", nil
}

func HandleAccess(ctx *context.Context) {
	redirect, userSocial, err := SocialAuth.OAuthAccess(ctx)
	if err != nil {
		beego.Error("SocialAuth.handleAccess", err)
	}

	if userSocial != nil {
		fmt.Sprintf("Identify: %s, AccessToken: %s", userSocial.Identify, userSocial.Data.AccessToken)
	}

	if len(redirect) > 0 {
		ctx.Redirect(302, redirect)
	}
}

func HandleRedirect(ctx *context.Context) {
	redirect, err := SocialAuth.OAuthRedirect(ctx)
	if err != nil {
		beego.Error("SocialAuth.handleRedirect", err)
	}

	if len(redirect) > 0 {
		ctx.Redirect(302, redirect)
	}
}