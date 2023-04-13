package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func regHello(context *gin.Context) {
	context.String(http.StatusOK, "hello")
}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", regHello)
	return router
}
