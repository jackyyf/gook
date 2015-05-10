package controllers

import (
	"github.com/astaxie/beego"
	"github.com/jackyyf/gook/models"
)

type IndexController struct {
	beego.Controller
}

func (this *IndexController) Get() {
	user, ok := this.Ctx.Input.Data["user"].(*models.User)
	if !ok {
		this.Abort("500")
	}
	this.Data["user"] = user
	this.TplNames = "index.tpl"
}
