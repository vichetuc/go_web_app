package guestbook

import (
	"html/template"
	"net/http"
	"time"

	"appengine"
	"appengine/datastore"
	"appengine/user"
	"github.com/juliofarah/go_web_app/api/models"
)

func init() {
	http.HandleFunc("/", root)
	http.HandleFunc("/sign", sign)
}

func guestbookKey(context appengine.Context) *datastore.Key {
	return datastore.NewKey(context, "Guestbook", "default_guestbook", 0, nil)
}

func root(w http.ResponseWriter, r *http.Request) {
	context := appengine.NewContext(r)
	query := datastore.NewQuery("Greeting").Ancestor(guestbookKey(context)).Order("-Date").Limit(10)
	greetings := make([]models.Greeting, 0, 10)
	if _, err := query.GetAll(context, &greetings); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := guestbookTemplate.Execute(w, greetings); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var guestbookTemplate = template.Must(template.New("book").Parse(`
<html>
  <head>
    <title>Go Guestbook</title>
  </head>
  <body>
    {{range .}}
      {{with .Author}}
        <p><b>{{.}}</b> wrote:</p>
      {{else}}
        <p>An anonymous person wrote:</p>
      {{end}}
      <pre>{{.Content}}</pre>
    {{end}}
    <form action="/sign" method="post">
      <div><textarea name="content" rows="3" cols="60"></textarea></div>
      <div><input type="submit" value="Sign Guestbook"></div>
    </form>
  </body>
</html>
`))

func sign(w http.ResponseWriter, r *http.Request) {
	context := appengine.NewContext(r)
	greeting := models.Greeting{
		Content: r.FormValue("content"),
		Date:    time.Now(),
	}
	if u := user.Current(context); u != nil {
		greeting.Author = u.String()
	}
	key := datastore.NewIncompleteKey(context, "Greeting", guestbookKey(context))
	_, err := datastore.Put(context, key, &greeting)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}