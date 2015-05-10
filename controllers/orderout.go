package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/jackyyf/gook/models"
	"strconv"
)

type OrderOutController struct {
	beego.Controller
}

func (c *OrderOutController) List() {
	c.TplNames = "orderout/list.tpl"
	errmsg, ok := c.GetSession("errmsg").(string)
	if ok && errmsg != "" {
		c.Data["errmsg"] = errmsg
	}
	c.DelSession("errmsg")
	orders, err := models.GetOrderOuts(nil)
	if err != nil {
		c.Data["content"] = err.Error()
		c.Abort("500")
	}
	c.Data["orders"] = orders
}

func (c *OrderOutController) Info() {
	c.TplNames = "orderout/info.tpl"
	sid, ok := c.Ctx.Input.Params["0"]
	id := 0
	var err error
	if !ok {
		c.Redirect("/orderout/list", 302)
		c.SetSession("errmsg", "Invalid order ID")
		return
	} else {
		id, err = strconv.Atoi(sid)
		if err != nil {
			c.Redirect("/orderout/list", 302)
			c.SetSession("errmsg", "Invalid order ID")
			return
		}
	}
	order, err := models.GetOrderOut(int32(id))
	if err != nil {
		c.Data["content"] = err.Error()
		c.Abort("500")
	}
	if order == nil {
		c.Redirect("/orderout/list", 302)
		c.SetSession("errmsg", fmt.Sprintf("Order ID %d does not exist!", id))
		return
	}
	errmsg, ok := c.GetSession("errmsg").(string)
	if ok && errmsg != "" {
		c.Data["errmsg"] = errmsg
	}
	c.DelSession("errmsg")
	c.Data["order"] = order
}
