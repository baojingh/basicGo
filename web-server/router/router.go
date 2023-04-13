package router

import (
	"basicGo/web-server/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

func regHello(context *gin.Context) {
	context.String(http.StatusOK, "hello")
}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/hello", regHello)
	router.GET("/user/:name", handler.UserSave)
	return router
}
