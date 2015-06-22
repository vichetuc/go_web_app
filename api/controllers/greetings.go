package controllers

import (
	"net/http"
	"appengine"
	"github.com/juliofarah/go_web_app/api/models"
	"html/template"
	"github.com/drborges/datastore-model"
	"log"
	"github.com/gin-gonic/gin"
)

//is there a way to make it better? Maybe parsing all the files isn't the most elegant approach
var templates = template.Must(template.ParseGlob("../api/views/*.html"))

func AllGreetings(c *gin.Context) {

	context := appengine.NewContext(c.Request)

	greetings := models.Greetings{}

	if err := db.NewDatastore(context).Query(greetings.GetAll()).All(&greetings); err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
	} else {
		printGreetings(greetings)
	}

	if err := templates.ExecuteTemplate(c.Writer, "guestbookPage", greetings); err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
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

func printGreetings(greetings models.Greetings) {
	for _, greeting := range greetings {
		log.Println("Author: " + greeting.Author)
		log.Println("Content: " + greeting.Content)
	}
}