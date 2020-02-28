package controller

import (
	"fmt"
	"gin-demo/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func AddPersonController(c *gin.Context) {
	var p model.Person
	firstName := c.Request.FormValue("firstName")
	lastName := c.Request.FormValue("lastName")
	p.FirstName = firstName
	p.LastName = lastName
	var id, err = p.AddPerson()
	if err != nil {
		log.Fatalln(err)
	}

	msg := fmt.Sprintf("insert person %d successful", id)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}

func GetAllPersonsController(context *gin.Context) {
	var person model.Person
	persons, _ := person.GetPersons()
	context.JSON(http.StatusOK, gin.H{
		"persons": persons,
	})
}

func GetPersonController(context *gin.Context) {
	var person model.Person
	cid := context.Param("id")
	person.Id, _ = strconv.Atoi(cid)

	person, err := person.GetPerson()
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{
			"person": nil,
		})
		return
	} else {
		context.JSON(http.StatusOK, gin.H{
			"person": person,
		})
	}
}

func UpdatePerson(context *gin.Context) {
	var p model.Person
	cid := context.Param("id")
	var firstName = context.Request.FormValue("firstName")
	var lastName = context.Request.FormValue("lastName")
	p.Id, _ = strconv.Atoi(cid)
	p.FirstName = firstName
	p.LastName = lastName
	p.UpdatePerson()
	msg := fmt.Sprintf("update person %d success", p.Id)
	context.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}

func DeletePerson(context *gin.Context) {
	var p model.Person

	cid := context.Param("id")
	p.Id, _ = strconv.Atoi(cid)
	p.DelPerson()

	msg := fmt.Sprintf("delete person %d success", p.Id)
	context.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}
