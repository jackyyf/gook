package routers

import (
	"github.com/astaxie/beego"
	c "github.com/jackyyf/gook/controllers"
)

func init() {
	beego.Router("/", &c.IndexController{})
	beego.AutoRouter(&c.UserController{})
}
