package controllers
import (
	"net/http"
	"appengine"
	"github.com/juliofarah/go_web_app/api/models"
	"time"
	"appengine/user"
	"appengine/datastore"
	"github.com/juliofarah/go_web_app/api/keys"
)

func Sign(w http.ResponseWriter, r *http.Request) {
	context := appengine.NewContext(r)
	greeting := models.Greeting{
		Content: r.FormValue("content"),
		Date:    time.Now(),
	}
	if u := user.Current(context); u != nil {
		greeting.Author = u.String()
	}
	key := datastore.NewIncompleteKey(context, "Greeting", keys.GuestbookKey(context))
	_, err := datastore.Put(context, key, &greeting)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

