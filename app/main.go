package guestbook

import (
	"net/http"
	"github.com/juliofarah/go_web_app/api/controllers"
)

func init() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/sign", controllers.Sign)
}