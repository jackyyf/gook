package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/jackyyf/gook/models"
	"strconv"
)

type OrderInController struct {
	beego.Controller
}

func (c *OrderInController) Prepare() {
	c.Data["status"] = models.Status
}

func (c *OrderInController) List() {
	c.TplNames = "orderin/list.tpl"
	errmsg, ok := c.GetSession("errmsg").(string)
	if ok && errmsg != "" {
		c.Data["errmsg"] = errmsg
	}
	c.DelSession("errmsg")
	orders, err := models.GetOrderIns(nil)
	if err != nil {
		c.Data["content"] = err.Error()
		c.Abort("500")
	}
	c.Data["orders"] = orders
}

func (c *OrderInController) Info() {
	c.TplNames = "orderin/info.tpl"
	sid, ok := c.Ctx.Input.Params["0"]
	id := 0
	var err error
	if !ok {
		c.Redirect("/orderin/list", 302)
		c.SetSession("errmsg", "Invalid order ID")
		return
	} else {
		id, err = strconv.Atoi(sid)
		if err != nil {
			c.Redirect("/orderin/list", 302)
			c.SetSession("errmsg", "Invalid order ID")
			return
		}
	}
	order, err := models.GetOrderIn(int32(id))
	if err != nil {
		c.Data["content"] = err.Error()
		c.Abort("500")
	}
	if order == nil {
		c.Redirect("/orderin/list", 302)
		c.SetSession("errmsg", fmt.Sprintf("Book ID %d does not exist!", id))
		return
	}
	errmsg, ok := c.GetSession("errmsg").(string)
	if ok && errmsg != "" {
		c.Data["errmsg"] = errmsg
	}
	c.DelSession("errmsg")
	c.Data["order"] = order
}

func (c *OrderInController) Pay() {
	c.TplNames = "orderin/info.tpl"
	sid, ok := c.Ctx.Input.Params["0"]
	id := 0
	var err error
	if !ok {
		c.Redirect("/orderin/list", 302)
		c.SetSession("errmsg", "Invalid order ID")
		return
	} else {
		id, err = strconv.Atoi(sid)
		if err != nil {
			c.Redirect("/orderin/list", 302)
			c.SetSession("errmsg", "Invalid order ID")
			return
		}
	}
	order, err := models.GetOrderIn(int32(id))
	if err != nil {
		c.Data["content"] = err.Error()
		c.Abort("500")
	}
	if order == nil {
		c.Redirect("/orderin/list", 302)
		c.SetSession("errmsg", fmt.Sprintf("Book ID %d does not exist!", id))
		return
	}
}
