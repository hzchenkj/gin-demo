package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexController(context *gin.Context) {
	context.JSON(http.StatusOK, "Hello world")
}
