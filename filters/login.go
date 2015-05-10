package filters

import (
	"fmt"
	"github.com/astaxie/beego/context"
	"github.com/jackyyf/gook/models"
	"github.com/jackyyf/gook/utils/log"
)

func LoginEnforcement(ctx *context.Context) {
	if ctx.Request.RequestURI == "/user/login" {
		return
	}
	uid, ok := ctx.Input.Session("user").(int32)
	if !ok {
		ctx.Input.CruSession.Set("errmsg", "Please login first.")
		log.Debug("Not login")
		ctx.Redirect(302, "/user/login")
		return
	}
	user, err := models.GetUser(uid)
	if err != nil {
		ctx.Abort(500, fmt.Sprintf("Error fetching user: %s", err))
	}
	if user == nil {
		ctx.Input.CruSession.Delete("user")
		ctx.Input.CruSession.Set("errmsg", "User was deleted.")
		ctx.Redirect(302, "/user/login")
	}
	ctx.Input.Data["user"] = user
	log.Debug("Logged in.")
}

func AdminProtection(ctx *context.Context) {
	uid, ok := ctx.Input.Session("user").(int32)
	if !ok {
		ctx.Redirect(302, "/")
		return
	}
	user, err := models.GetUser(uid)
	if err != nil || user == nil || !user.IsAdmin() {
		ctx.Redirect(302, "/")
	}
}
