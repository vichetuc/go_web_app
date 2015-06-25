package controllers

import (
	"net/http"
	"appengine"
	"github.com/juliofarah/go_web_app/api/models"
	"html/template"
	"github.com/drborges/datastore-model"
	"github.com/gin-gonic/gin"
)

//is there a way to make it better? Maybe parsing all the files isn't the most elegant approach
var templates = template.Must(template.ParseGlob("../api/views/*.html"))

func AllGreetings(c *gin.Context) {

	context := appengine.NewContext(c.Request)
	greetings := models.Greetings{}

	if err := db.NewDatastore(context).Query(greetings.GetAll()).All(&greetings); err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
	}

	if err := templates.ExecuteTemplate(c.Writer, "guestbookPage", greetings); err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
	}
}

func New(c *gin.Context) {

	context := appengine.NewContext(c.Request)

	greetings := models.Greetings{}
	//how to fix it?
	//understand whats the difference between calling a method that returns a Greeting
	//and creating a models.Greeting{}
	greeting := greetings.New("", "")

	c.BindJSON(&greeting)

	if err := db.NewDatastore(context).Create(greeting); err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GreetingsToJson(c *gin.Context) {

	context := appengine.NewContext(c.Request)
	greetings := models.Greetings{}

	if err := db.NewDatastore(context).Query(greetings.GetAll()).All(&greetings); err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
	} else {
		c.JSON(http.StatusOK, greetings)
	}
}
