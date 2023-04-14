package router

import (
	"basicGo/web-server/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.StaticFile("/datou.jpg", "/datou.jpg")

	router.GET("/hello", handler.Hello)

	userRouter := router.Group("/user")
	{
		userRouter.GET("/:name", handler.UserSave)
		userRouter.GET("", handler.UserSaveByQuery)
		userRouter.POST("/register", handler.UserRegister)
	}

	return router
}
