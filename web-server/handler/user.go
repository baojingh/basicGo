package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserSave(context *gin.Context) {
	username := context.Param("name")
	fmt.Println("***********", username)
	context.String(http.StatusOK, "saved")
}
