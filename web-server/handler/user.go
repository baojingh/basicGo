package handler

import (
	"basicGo/web-server/model"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserSave(context *gin.Context) {
	username := context.Param("name")
	fmt.Println("***********", username)
	context.String(http.StatusOK, "saved")
}

func UserSaveByQuery(context *gin.Context) {
	username := context.Query("name")
	age := context.DefaultQuery("age", "33")
	fmt.Println(username, age)
	context.String(http.StatusOK, "saved")
}

// func UserRegister(context *gin.Context) {
// 	var user model.Usermodel
// 	if err := context.ShouldBind(&user); err != nil {
// 		fmt.Println("err ->", err.Error())
// 		return
// 	}
// 	log.Println("#######", user.Email)
// 	log.Println("#######", user.Password)
// 	id := user.Save()
// 	log.Println("#######", id)
// 	context.Redirect(http.StatusMovedPermanently, "/")
// }

func UserRegister(context *gin.Context) {
	user := model.User{}
	if e := context.ShouldBind(&user); e == nil {
		user.Save()
		log.Println("user is ", user)
	}
	context.String(http.StatusOK, "result")
}

func Hello(context *gin.Context) {
	context.String(http.StatusOK, "hello")
}
