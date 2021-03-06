package socialAuth

import (
	"fmt"
	"github.com/EPICPaaS/social-auth"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"strconv"
)

var SocialAuth *social.SocialAuth

type SocialAuther struct {
}

func (p *SocialAuther) IsUserLogin(ctx *context.Context) (int, bool) {
	if id, ok := ctx.Input.CruSession.Get("login_user").(int); ok && id == 1 {
		return id, true
	}
	return 0, false
}

func (p *SocialAuther) LoginUser(ctx *context.Context, uid int) (string, error) {
	fmt.Println("uid--" + strconv.Itoa(uid))
	// fake login the user
	ctx.Input.CruSession.Set("login_user", uid)
	return "/register/connect", nil
}

func HandleAccess(ctx *context.Context) {
	redirect, userSocial, err := SocialAuth.OAuthAccess(ctx)
	if err != nil {
		beego.Error("SocialAuth.handleAccess", err)
	}

	if userSocial != nil {
		fmt.Println("Identify: %s, AccessToken: %s", userSocial.Identify, userSocial.Data.AccessToken)
	}
	ctx.Input.CruSession.Set("custom_userSocial_identify", userSocial.Identify)
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
