package router

import (
	. "gin-demo/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", IndexController)
	router.POST("/person/add", AddPersonController)
	router.GET("/persons", GetAllPersonsController)
	router.GET("/person/:id", GetPersonController)
	router.DELETE("/person/:id", DeletePerson)
	router.PUT("/person/:id", UpdatePerson)
	return router
}
