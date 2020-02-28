package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexController(context *gin.Context) {
	context.String(http.StatusOK, "Hello Gin World")
	//context.JSON(http.StatusOK, "Hello World")
}
