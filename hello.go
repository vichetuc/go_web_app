package go_web_app
import (
	"html/template"
	"io/ioutil"
	"net/http"
)

var (
	guestbookForm []byte
	signTemplate = template.Must(template.ParseFiles("views/guestbook.html"))
)

func init() {
	content, err := ioutil.ReadFile("views/guestbookform.html")

	if err != nil {
		panic(err)
	}

	guestbookForm = content

	http.HandleFunc("/", root)
	http.HandleFunc("/sign", sign)
}

func root(w http.ResponseWriter, r *http.Request) { w.Write(guestbookForm) }

func sign(w http.ResponseWriter, r *http.Request) {
	err := signTemplate.Execute(w, r.FormValue("content"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
