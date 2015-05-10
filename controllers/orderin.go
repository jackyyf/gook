package controllers

import (
	// "fmt"
	"github.com/astaxie/beego"
	"github.com/jackyyf/gook/models"
	// "strconv"
)

type OrderInController struct {
	beego.Controller
}

func (c *OrderInController) List() {
	c.TplNames = "orderin/list.tpl"
	errmsg, ok := c.GetSession("errmsg").(string)
	if ok && errmsg != "" {
		c.Data["errmsg"] = errmsg
	}
	c.DelSession("errmsg")
	books, err := models.SearchBooks("", nil, nil, nil, -1, -1)
	if err != nil {
		c.Data["content"] = err.Error()
		c.Abort("500")
	}
	c.Data["books"] = books
}
