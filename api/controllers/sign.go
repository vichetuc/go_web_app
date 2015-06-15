package controllers
import (
	"net/http"
	"appengine"
	"github.com/juliofarah/go_web_app/api/models"
	"time"
	"appengine/user"
	"github.com/drborges/datastore-model"
)

func Sign(w http.ResponseWriter, r *http.Request) {

	context := appengine.NewContext(r)

	greeting := new(models.Greeting)
	greeting.Content = r.FormValue("content")
	greeting.Date = time.Now()

	if u := user.Current(context); u != nil {
		greeting.Author = u.String()
	}

	if err := db.NewDatastore(context).Create(greeting); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

