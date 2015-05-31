package main

import (
	"net/http"
	"log"
	"path/filepath"
	"os"
	"bufio"
)

var contentTypesMap = map[string]string{
	".css" : "text/css",
	".html" : "text/html",
	".js" : "application/javascript",
	".png" : "image/png",
	".svg" : "image/svg+xml",
	".mp4" : "video/mp4",
}

type MyHandler struct {
}

func (this *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]
	log.Println(path)

	file, err := os.Open(path)

	if err == nil {
		bufferedReader := bufio.NewReader(file)
		contentType, present := contentTypesMap[filepath.Ext(path)]
		if !present { contentType = "plain/text" }

		w.Header().Add("Content Type", contentType)
		bufferedReader.WriteTo(w)
	} else {
		w.WriteHeader(404)
		w.Write([]byte("404 my friend - " + http.StatusText(404)))
	}
}

func main() {
	http.Handle("/", new(MyHandler))
	http.ListenAndServe(":8080", nil)
}