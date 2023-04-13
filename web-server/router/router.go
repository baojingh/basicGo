package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", func(contenxt *gin.Context) {
		contenxt.String(http.StatusOK, "hello")
	})
	return router
}
