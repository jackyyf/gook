package controllers

import (
	"github.com/astaxie/beego"
	"github.com/jackyyf/gook/models"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) Login() {
	if this.Ctx.Input.GetData("user") != nil {
		this.Redirect("/", 302)
		return
	}
	if this.Ctx.Request.Method == "GET" {
		errmsg, ok := this.GetSession("errmsg").(string)
		if ok && errmsg != "" {
			this.Data["errmsg"] = errmsg
		}
		this.DelSession("errmsg")
		this.TplNames = "user/login.tpl"
	} else if this.Ctx.Request.Method == "POST" {
		var username, password string
		this.Ctx.Input.Bind(&username, "username")
		this.Ctx.Input.Bind(&password, "password")
		user, err := models.GetUserByName(username)
		if err != nil {
			this.Data["content"] = err.Error()
			this.Abort("500")
		}
		if user == nil {
			this.Data["errmsg"] = "Invalid username or password."
			this.TplNames = "user/login.tpl"
			return
		}
	}
}
