package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/jackyyf/gook/models"
	"time"
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
		username := this.GetString("username")
		password := this.GetString("password")
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
		if user.Login(password) {
			this.SetSession("user", user.ID())
			this.Redirect("/", 302)
			return
		} else {
			this.Data["errmsg"] = "Invalid username or password."
			this.TplNames = "user/login.tpl"
			return
		}
	}
}

func (this *UserController) Logout() {
	this.DelSession("user")
	this.Redirect("/user/login", 302)
}

func (this *UserController) Me() {
	user, ok := this.Ctx.Input.Data["user"].(*models.User)
	if !ok {
		this.Abort("500")
	}
	this.TplNames = "user/me.tpl"
	if this.Ctx.Request.Method == "GET" {
		errmsg, ok := this.GetSession("errmsg").(string)
		if ok && errmsg != "" {
			this.Data["errmsg"] = errmsg
		}
		this.DelSession("errmsg")
	} else if this.Ctx.Request.Method == "POST" {
		password := this.GetString("password")
		realname := this.GetString("realname")
		born := this.GetString("born")
		gender, err := this.GetInt("gender")
		if err != nil {
			this.Data["errmsg"] = fmt.Sprintf("Invalid gender: %s", err.Error())
			return
		}
		npass := this.GetString("npass")
		nrpass := this.GetString("nrpass")
		if !user.Login(password) {
			this.Data["errmsg"] = "Incorrect password"
			return
		}
		if realname != "" {
			user.RealName = realname
		}
		if born != "" {
			btime, err := time.Parse("2006-01-02", born)
			if err != nil {
				this.Data["errmsg"] = fmt.Sprintf("Invalid born: %s", err.Error())
				return
			}
			user.Born = btime
		}
		if gender != 0 && gender != 1 {
			this.Data["errmsg"] = fmt.Sprintf("Invalid gender: %d", gender)
			return
		}
		user.Gender = int32(gender)
		if npass != nrpass {
			this.Data["errmsg"] = "New password doesn't match"
			return
		}
		if npass != "" {
			user.SetPassword(npass)
		}
		user.Save()
	}
}
