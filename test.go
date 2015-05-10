package main

import (
	"time"
	. "github.com/jackyyf/gook/models"
)

func main() {
	user := new(User)
	user.SetPassword("123456")
	user.SetAdmin()
	user.Name = "jackyyf"
	user.RealName = "余一夫"
	user.Gender = 1
	user.Born = time.Date(1995, time.July, 3, 0, 0, 0, 0, time.UTC)
	user.Create()
}
