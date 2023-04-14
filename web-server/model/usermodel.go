package model

import (
	"basicGo/web-server/initDb"
	"log"
)

type Usermodel struct {
	Email    string `form:"email" binding:"email"`
	Password string `form:"password" binding: "password"`
}

type User struct {
	Id       int    `form:"id" 		example:"1"`
	Email    string `form:"email" 	 example:"js"`
	Password string `form:"password" example: "123"`
}

func (user User) Save() int {
	log.Println("start save *************")
	res := initDb.Db.Create(&user)
	log.Println("finish save *************")
	if res.Error != nil {
		return 0
	}
	return 1

}
