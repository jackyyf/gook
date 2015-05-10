package main

import (
	"github.com/astaxie/beego"
	_ "github.com/jackyyf/gook/db"
	. "github.com/jackyyf/gook/filters"
	_ "github.com/jackyyf/gook/routers"
)

func main() {
	beego.InsertFilter("", beego.BeforeRouter, LoginEnforcement)
	beego.InsertFilter("*", beego.BeforeRouter, LoginEnforcement)
	beego.InsertFilter("/admin*", beego.BeforeRouter, AdminProtection)
	beego.Run()
}
