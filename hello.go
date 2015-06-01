package go_web_app

import (
	"net/http"
	"fmt"
	"html/template"
)

func init() {
	http.HandleFunc("/", root)
	http.HandleFunc("/sign", sign)

}

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, guestbookForm)
}

const guestbookForm = `
	<html>
		<body>
			<form actions="/sign" method="post">
				<div><textarea name="content" rows="3" cols="60"></textarea></div>
				<div><input type="submit" value="Sign Guestbook"></div>
			</form>
		</body>
	</html>
`

func sign(w http.ResponseWriter, r *http.Request) {
	err := signTemplate.Execute(w, r.FormValue("content"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

const signTemplateHTML = `
	<html>
		<body>
			<p>You wrote:</p>
			<pre>{{.}}</pre>
		</body>
	</html>
`

var signTemplate = template.Must(template.New("sign").Parse(signTemplateHTML))