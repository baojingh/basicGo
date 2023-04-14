package model

import (
	"basicGo/web-server/initDb"
	"log"
)

type Usermodel struct {
	Email    string `form:"email" binding:"email"`
	Password string `form:"password" binding: "password"`
}

func (user *Usermodel) Save() int64 {
	log.Println("start save *************")

	result, err := initDb.Db.Exec("insert into user(email, password) values (?, ?)", user.Email, user.Password)
	if err != nil {
		log.Panicln("user insert ")
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Panicln("insert err, cannot get id")
	}
	log.Println("finish save *************")

	return id
}
