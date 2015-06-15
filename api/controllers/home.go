package controllers

import (
	"net/http"
	"appengine"
	"github.com/juliofarah/go_web_app/api/models"
	"html/template"
	"github.com/drborges/datastore-model"
	"log"
)

//is there a way to make it better? Maybe parsing all the files isn't the most elegant approach
var templates = template.Must(template.ParseGlob("../api/views/*"))

func Home(w http.ResponseWriter, r *http.Request) {

	context := appengine.NewContext(r)

	greetings := models.Greetings{}

	if err := db.NewDatastore(context).Query(greetings.GetAll()).All(&greetings); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		printGreetings(greetings)
	}

	if err := templates.ExecuteTemplate(w, "guestbookPage", greetings); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func printGreetings(greetings models.Greetings) {
	for _, greeting := range greetings {
		log.Println("Author: " + greeting.Author)
		log.Println("Content: " + greeting.Content)
	}
}