package controllers

import (
	"net/http"
	"appengine"
	"github.com/juliofarah/go_web_app/api/models"
	"github.com/drborges/datastore-model"
	"github.com/gin-gonic/gin"
	"github.com/juliofarah/go_web_app/api/modelAndView"
)

func Create(c *gin.Context) {
	//can I extract context and datastore ?
	context := appengine.NewContext(c.Request)
	greetingParams := modelAndView.GreetingAsForm{}

	c.Bind(&greetingParams)

	//understand the difference between new(Greeting)
	//and variable := Greeting{}
	greetings := models.Greetings{}
	greeting := greetings.New(greetingParams.Author, greetingParams.Content)

	if err := db.NewDatastore(context).Create(greeting); err != nil {
		statusInternalServerError(c, err)
		return
	}
	c.String(201, "")

}

func List(c *gin.Context) {

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
