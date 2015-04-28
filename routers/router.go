package routers

import (
	"github.com/jackyyf/gook/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
