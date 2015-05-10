package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/jackyyf/gook/models"
	"github.com/jackyyf/gook/utils/isbn"
	"strconv"
)

type BookController struct {
	beego.Controller
}

const PAGE_SIZE int = 30

func (c *BookController) List() {
	c.TplNames = "book/list.tpl"
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

func (c *BookController) New() {
	if c.Ctx.Request.Method == "GET" {
		errmsg, ok := c.GetSession("errmsg").(string)
		if ok && errmsg != "" {
			c.Data["errmsg"] = errmsg
		}
		c.DelSession("errmsg")
		c.TplNames = "book/new.tpl"
	} else if c.Ctx.Request.Method == "POST" {
		book := new(models.Book)
		title := c.GetString("title")
		if title == "" {
			c.SetSession("errmsg", "Title can't be empty")
			c.Redirect("/book/new", 302)
			return
		}
		book.Name = title
		publisher := c.GetString("publisher")
		if title == "" {
			c.SetSession("errmsg", "Publisher can't be empty")
			c.Redirect("/book/new", 302)
			return
		}
		book.Publisher = publisher
		author := c.GetString("author")
		if title == "" {
			c.SetSession("errmsg", "Author can't be empty")
			c.Redirect("/book/new", 302)
			return
		}
		book.Author = author
		price, err := c.GetFloat("price")
		if err != nil {
			c.SetSession("errmsg", "Invalid price: "+err.Error())
			c.Redirect("/book/new", 302)
			return
		}
		book.Price = price
		Isbn := isbn.Normalize(c.GetString("isbn"))
		if Isbn == "" {
			c.SetSession("errmsg", "Invalid ISBN")
			c.Redirect("/book/new", 302)
			return
		}
		book.ISBN = Isbn
		err = book.Create()
		if err != nil {
			c.SetSession("errmsg", err.Error())
			c.Redirect("/book/new", 302)
			return
		}
		c.Redirect(fmt.Sprintf("/book/info/%d", book.ID()), 302)
	}
}

func (c *BookController) Info() {
	c.TplNames = "book/info.tpl"
	sid, ok := c.Ctx.Input.Params["0"]
	id := 0
	var err error
	if !ok {
		c.Redirect("/book/list", 302)
		c.SetSession("errmsg", "Invalid Book ID")
		return
	} else {
		id, err = strconv.Atoi(sid)
		if err != nil {
			c.Redirect("/book/list", 302)
			c.SetSession("errmsg", "Invalid Book ID")
			return
		}
	}
	book, err := models.GetBook(int32(id))
	if err != nil {
		c.Data["content"] = err.Error()
		c.Abort("500")
	}
	if book == nil {
		c.Redirect("/book/list", 302)
		c.SetSession("errmsg", fmt.Sprintf("Book ID %d does not exist!", id))
		return
	}
	c.Data["book"] = book

}

func (c *BookController) Buy() {
	sid, ok := c.Ctx.Input.Params["0"]
	id := 0
	var err error
	if !ok {
		c.Redirect("/book/list", 302)
		c.SetSession("errmsg", "Invalid Book ID")
		return
	} else {
		id, err = strconv.Atoi(sid)
		if err != nil {
			c.Redirect("/book/list", 302)
			c.SetSession("errmsg", "Invalid Book ID")
			return
		}
	}
	book, err := models.GetBook(int32(id))
	if err != nil {
		c.Data["content"] = err.Error()
		c.Abort("500")
	}
	if book == nil {
		c.Redirect("/book/list", 302)
		c.SetSession("errmsg", fmt.Sprintf("Book ID %d does not exist!", id))
		return
	}
	if c.Ctx.Request.Method == "GET" {
		errmsg, ok := c.GetSession("errmsg").(string)
		if ok && errmsg != "" {
			c.Data["errmsg"] = errmsg
		}
		c.DelSession("errmsg")
		c.Data["book"] = book
		c.TplNames = "book/buy.tpl"
	} else if c.Ctx.Request.Method == "POST" {
		order := new(models.OrderIn)
		order.Book = book
		price, err := c.GetFloat("price")
		if err != nil {
			c.SetSession("errmsg", "Invalid price: "+err.Error())
			c.Redirect(fmt.Sprintf("/book/buy/%d", id), 302)
			return
		}
		order.Price = price
		amount, err := c.GetInt32("amount")
		if err != nil {
			c.SetSession("errmsg", "Invalid amount: "+err.Error())
			c.Redirect(fmt.Sprintf("/book/buy/%d", id), 302)
			return
		}
		if amount <= 0 {
			c.SetSession("errmsg", fmt.Sprintf("Invalid amount: %d <= 0", amount))
			c.Redirect(fmt.Sprintf("/book/buy/%d", id), 302)
			return
		}
		order.Amount = amount
		err = order.Create()
		if err != nil {
			c.SetSession("errmsg", err.Error())
			c.Redirect(fmt.Sprintf("/book/buy/%d", id), 302)
			return
		}
		c.Redirect(fmt.Sprintf("/orderin/info/%d", order.ID()), 302)
	}
}

func (c *BookController) Sell() {
	sid, ok := c.Ctx.Input.Params["0"]
	id := 0
	var err error
	if !ok {
		c.Redirect("/book/list", 302)
		c.SetSession("errmsg", "Invalid Book ID")
		return
	} else {
		id, err = strconv.Atoi(sid)
		if err != nil {
			c.Redirect("/book/list", 302)
			c.SetSession("errmsg", "Invalid Book ID")
			return
		}
	}
	book, err := models.GetBook(int32(id))
	if err != nil {
		c.Data["content"] = err.Error()
		c.Abort("500")
	}
	if book == nil {
		c.Redirect("/book/list", 302)
		c.SetSession("errmsg", fmt.Sprintf("Book ID %d does not exist!", id))
		return
	}
	if c.Ctx.Request.Method == "GET" {
		errmsg, ok := c.GetSession("errmsg").(string)
		if ok && errmsg != "" {
			c.Data["errmsg"] = errmsg
		}
		c.DelSession("errmsg")
		c.Data["book"] = book
		c.TplNames = "book/sell.tpl"
	} else if c.Ctx.Request.Method == "POST" {
		order := new(models.OrderOut)
		order.Book = book
		amount, err := c.GetInt32("amount")
		if err != nil {
			c.SetSession("errmsg", "Invalid amount: "+err.Error())
			c.Redirect(fmt.Sprintf("/book/sell/%d", id), 302)
			return
		}
		if amount <= 0 {
			c.SetSession("errmsg", fmt.Sprintf("Invalid amount: %d <= 0", amount))
			c.Redirect(fmt.Sprintf("/book/sell/%d", id), 302)
			return
		}
		if amount > book.Amount {
			c.SetSession("errmsg", fmt.Sprintf("Out of stock: only %d books left", book.Amount))
			c.Redirect(fmt.Sprintf("/book/sell/%d", id), 302)
			return
		}
		order.Amount = amount
		err = order.Create()
		if err != nil {
			c.SetSession("errmsg", err.Error())
			c.Redirect(fmt.Sprintf("/book/sell/%d", id), 302)
			return
		}
		c.Redirect(fmt.Sprintf("/orderout/info/%d", order.ID()), 302)
	}
}

func (c *BookController) Edit() {
	sid, ok := c.Ctx.Input.Params["0"]
	id := 0
	var err error
	if !ok {
		c.Redirect("/book/list", 302)
		c.SetSession("errmsg", "Invalid Book ID")
		return
	} else {
		id, err = strconv.Atoi(sid)
		if err != nil {
			c.Redirect("/book/list", 302)
			c.SetSession("errmsg", "Invalid Book ID")
			return
		}
	}
	book, err := models.GetBook(int32(id))
	if err != nil {
		c.Data["content"] = err.Error()
		c.Abort("500")
	}
	if book == nil {
		c.Redirect("/book/list", 302)
		c.SetSession("errmsg", fmt.Sprintf("Book ID %d does not exist!", id))
		return
	}
	title := c.GetString("title")
	if title == "" {
		c.SetSession("errmsg", "Title can't be empty")
		c.Redirect(fmt.Sprintf("/book/info/%d", id), 302)
		return
	}
	book.Name = title
	publisher := c.GetString("publisher")
	if title == "" {
		c.SetSession("errmsg", "Publisher can't be empty")
		c.Redirect(fmt.Sprintf("/book/info/%d", id), 302)
		return
	}
	book.Publisher = publisher
	author := c.GetString("author")
	if title == "" {
		c.SetSession("errmsg", "Author can't be empty")
		c.Redirect(fmt.Sprintf("/book/info/%d", id), 302)
		return
	}
	book.Author = author
	price, err := c.GetFloat("price")
	if err != nil {
		c.SetSession("errmsg", "Invalid price: "+err.Error())
		c.Redirect(fmt.Sprintf("/book/info/%d", id), 302)
		return
	}
	book.Price = price
	Isbn := isbn.Normalize(c.GetString("isbn"))
	if Isbn == "" {
		c.SetSession("errmsg", "Invalid ISBN")
		c.Redirect("/book/new", 302)
		return
	}
	book.ISBN = Isbn
	err = book.Save()
	if err != nil {
		c.SetSession("errmsg", err.Error())
		c.Redirect(fmt.Sprintf("/book/info/%d", id), 302)
		return
	}
	c.Redirect(fmt.Sprintf("/book/info/%d", id), 302)
}

/*

func (c *BookController) Delete() {
	sid, ok := c.Ctx.Input.Params["0"]
	id := 0
	var err error
	if !ok {
		c.Redirect("/book/list", 302)
		c.SetSession("errmsg", "Invalid Book ID")
		return
	} else {
		id, err = strconv.Atoi(sid)
		if err != nil {
			c.Redirect("/book/list", 302)
			c.SetSession("errmsg", "Invalid Book ID")
			return
		}
	}
	book, err := models.GetBook(int32(id))
	if err != nil {
		c.Data["content"] = err.Error()
		c.Abort("500")
	}
	if book == nil {
		c.Redirect("/book/list", 302)
		c.SetSession("errmsg", fmt.Sprintf("Book ID %d does not exist!", id))
		return
	}
	err = book.Delete()
	if err != nil {
		c.SetSession("errmsg", err.Error())
	}
	c.Redirect("/book/list", 302)
}

*/
