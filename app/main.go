package guestbook

import (
	"html/template"
	"net/http"

	"appengine"
	"appengine/datastore"
	"github.com/juliofarah/go_web_app/api/models"
	"github.com/juliofarah/go_web_app/api/controllers"
)

func init() {
	http.HandleFunc("/", root)
	http.HandleFunc("/sign", controllers.Sign)
}

func root(w http.ResponseWriter, r *http.Request) {
	context := appengine.NewContext(r)
	query := datastore.NewQuery("Greeting").Ancestor(controllers.GuestbookKey(context)).Order("-Date").Limit(10)
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