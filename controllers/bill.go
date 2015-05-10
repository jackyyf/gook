package controllers

import (
	"github.com/astaxie/beego"
	"github.com/jackyyf/gook/models"
	"github.com/jackyyf/gook/utils/log"
	"time"
)

type BillController struct {
	beego.Controller
}

func (c *BillController) List() {
	c.TplNames = "bill/list.tpl"
	errmsg, ok := c.GetSession("errmsg").(string)
	if ok && errmsg != "" {
		c.Data["errmsg"] = errmsg
	}
	c.DelSession("errmsg")
	t := time.Time{}
	aft := c.GetString("after")
	t, err := time.Parse("2006-01-02 15:04:05 -0700", aft+" +0800")
	if err != nil {
		log.Warn("Invalid date %s: %s", aft, err)
		t = time.Time{}
	} else {
		c.Data["after"] = t
	}
	bills, err := models.GetBillsAfter(t)
	if err != nil {
		c.Data["content"] = err.Error()
		c.Abort("500")
	}

	c.Data["bills"] = bills
}
