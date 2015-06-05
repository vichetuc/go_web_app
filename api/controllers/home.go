package controllers

import (
	"net/http"
	"appengine"
	"appengine/datastore"
	"github.com/juliofarah/go_web_app/api/keys"
	"github.com/juliofarah/go_web_app/api/models"
	"html/template"
)

var templates = template.Must(template.ParseGlob("../api/views/*"))

func Home(w http.ResponseWriter, r *http.Request) {
  context := appengine.NewContext(r)
  query := datastore.NewQuery("Greeting").Ancestor(keys.GuestbookKey(context)).Order("-Date").Limit(10)
  greetings := make([]models.Greeting, 0, 10)
  if _, err := query.GetAll(context, &greetings); err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
  if err := templates.ExecuteTemplate(w, "guestbookPage", greetings); err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}