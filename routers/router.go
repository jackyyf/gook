package routers

import (
	"github.com/astaxie/beego"
	c "github.com/jackyyf/gook/controllers"
)

func init() {
	beego.Router("/", &c.IndexController{})
	beego.AutoRouter(&c.UserController{})
	beego.AutoRouter(&c.BookController{})
	beego.AutoRouter(&c.OrderInController{})
	beego.AutoRouter(&c.OrderOutController{})
	beego.AutoRouter(&c.BillController{})
	beego.AutoRouter(&c.AdminController{})
}
