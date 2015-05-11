package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/jackyyf/gook/models"
	"strconv"
	"time"
)

type AdminController struct {
	beego.Controller
}

func (c *AdminController) Prepare() {
	user, ok := c.Ctx.Input.Data["user"].(*models.User)
	if !ok {
		c.Abort("500")
	}
	if user == nil || !user.IsAdmin() {
		c.Redirect("/", 302)
		c.StopRun()
	}
}

func (c *AdminController) List() {
	c.TplNames = "admin/list.tpl"
	errmsg, ok := c.GetSession("errmsg").(string)
	if ok && errmsg != "" {
		c.Data["errmsg"] = errmsg
	}
	c.DelSession("errmsg")
	users, err := models.GetUsers()
	if err != nil {
		c.Data["content"] = err.Error()
		c.Abort("500")
	}
	c.Data["users"] = users
}

func (this *AdminController) Info() {
	sid, ok := this.Ctx.Input.Params["0"]
	id := 0
	var err error
	if !ok {
		this.Redirect("/admin/list", 302)
		this.SetSession("errmsg", "Invalid user ID")
		return
	} else {
		id, err = strconv.Atoi(sid)
		if err != nil {
			this.Redirect("/admin/list", 302)
			this.SetSession("errmsg", "Invalid user ID")
			return
		}
	}
	nuser, err := models.GetUser(int32(id))
	this.TplNames = "admin/info.tpl"
	if this.Ctx.Request.Method == "GET" {
		errmsg, ok := this.GetSession("errmsg").(string)
		if ok && errmsg != "" {
			this.Data["errmsg"] = errmsg
		}
		this.Data["nuser"] = nuser
		this.DelSession("errmsg")
	} else if this.Ctx.Request.Method == "POST" {
		password := this.GetString("password")
		realname := this.GetString("realname")
		if password != "" {
			nuser.SetPassword(password)
		}
		born := this.GetString("born")
		gender, err := this.GetInt("gender")
		if err != nil {
			this.Data["errmsg"] = fmt.Sprintf("Invalid gender: %s", err.Error())
			return
		}
		admin, err := this.GetInt("admin")
		if err != nil {
			this.Data["errmsg"] = fmt.Sprintf("Invalid admin: %s", err.Error())
			return
		}
		if realname != "" {
			nuser.RealName = realname
		}
		if born != "" {
			btime, err := time.Parse("2006-01-02", born)
			if err != nil {
				this.Data["errmsg"] = fmt.Sprintf("Invalid born: %s", err.Error())
				return
			}
			nuser.Born = btime
		}
		if gender != 0 && gender != 1 {
			this.Data["errmsg"] = fmt.Sprintf("Invalid gender: %d", gender)
			return
		}
		nuser.Gender = int32(gender)
		if admin != 0 && admin != 1 {
			this.Data["errmsg"] = fmt.Sprintf("Invalid admin: %d", admin)
			return
		}
		if admin == 0 {
			nuser.ClearAdmin()
		} else if admin == 1 {
			nuser.SetAdmin()
		}
		nuser.Gender = int32(gender)
		nuser.Save()
		this.Redirect(fmt.Sprintf("/admin/info/%d", nuser.ID()), 302)
	}
}

func (this *AdminController) New() {
	this.TplNames = "admin/new.tpl"
	if this.Ctx.Request.Method == "GET" {
		errmsg, ok := this.GetSession("errmsg").(string)
		if ok && errmsg != "" {
			this.Data["errmsg"] = errmsg
		}
		this.DelSession("errmsg")
	} else if this.Ctx.Request.Method == "POST" {
		nuser := new(models.User)
		username := this.GetString("username")
		password := this.GetString("password")
		realname := this.GetString("realname")
		if username == "" {
			this.Data["errmsg"] = "Username required"
			return
		}
		nuser.Name = username
		if password == "" {
			this.Data["errmsg"] = "Password required"
			return
		}
		if realname == "" {
			this.Data["errmsg"] = "Realname required"
			return
		}
		nuser.SetPassword(password)
		born := this.GetString("born")
		gender, err := this.GetInt("gender")
		if err != nil {
			this.Data["errmsg"] = fmt.Sprintf("Invalid gender: %s", err.Error())
			return
		}
		admin, err := this.GetInt("admin")
		if err != nil {
			this.Data["errmsg"] = fmt.Sprintf("Invalid admin: %s", err.Error())
			return
		}
		nuser.RealName = realname
		if born != "" {
			btime, err := time.Parse("2006-01-02", born)
			if err != nil {
				this.Data["errmsg"] = fmt.Sprintf("Invalid born: %s", err.Error())
				return
			}
			nuser.Born = btime
		} else {
			this.Data["errmsg"] = "Born date required"
			return
		}
		if gender != 0 && gender != 1 {
			this.Data["errmsg"] = fmt.Sprintf("Invalid gender: %d", gender)
			return
		}
		nuser.Gender = int32(gender)
		if admin != 0 && admin != 1 {
			this.Data["errmsg"] = fmt.Sprintf("Invalid admin: %d", admin)
			return
		}
		if admin == 0 {
			nuser.ClearAdmin()
		} else if admin == 1 {
			nuser.SetAdmin()
		}
		nuser.Gender = int32(gender)
		nuser.Create()
		this.Redirect(fmt.Sprintf("/admin/info/%d", nuser.ID()), 302)
	}
}

func (c *AdminController) Remove() {
	sid, ok := c.Ctx.Input.Params["0"]
	id := 0
	var err error
	if !ok {
		c.Redirect("/admin/list", 302)
		c.SetSession("errmsg", "Invalid user ID")
		return
	} else {
		id, err = strconv.Atoi(sid)
		if err != nil {
			c.Redirect("/admin/list", 302)
			c.SetSession("errmsg", "Invalid user ID")
			return
		}
	}
	nuser, err := models.GetUser(int32(id))
	if err != nil {
		c.Data["content"] = err.Error()
		c.Abort("500")
	}
	if nuser == nil {
		c.Redirect("/admin/list", 302)
		c.SetSession("errmsg", fmt.Sprintf("User ID %d does not exist!", id))
		return
	}
	err = nuser.Delete()
	if err != nil {
		c.SetSession("errmsg", err.Error())
	}
	c.Redirect("/admin/list", 302)
}
