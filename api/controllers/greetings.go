package controllers

import (
	"net/http"
	"appengine"
	"github.com/juliofarah/go_web_app/api/models"
	"github.com/drborges/datastore-model"
	"github.com/gin-gonic/gin"
)

func New(c *gin.Context) {
	context := appengine.NewContext(c.Request)

	//how to fix it?
	//understand whats the difference between calling a method that returns a Greeting
	//and creating a models.Greeting{}

	greetingParams := models.GreetingAsForm{}

	c.Bind(&greetingParams)

	greetings := models.Greetings{}
	greeting := greetings.New(greetingParams.Author, greetingParams.Content)

	if err := db.NewDatastore(context).Create(greeting); err != nil {
		statusInternalServerError(c, err)
		return
	}
	c.String(201, "")

}

func AllGreetings(c *gin.Context) {

	context := appengine.NewContext(c.Request)
	greetings := models.Greetings{}

	if err := db.NewDatastore(context).Query(greetings.GetAll()).All(&greetings); err != nil {
		statusInternalServerError(c, err)
	} else {
		c.JSON(http.StatusOK, greetings)
	}
}

func statusInternalServerError(c *gin.Context, err error) {
	http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
}
