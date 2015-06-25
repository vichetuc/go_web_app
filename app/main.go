package guestbook

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/juliofarah/go_web_app/api/controllers"
)

func init() {

	r := gin.New()

	r.GET("/greetings/all", controllers.GreetingsToJson)
	r.POST("/greetings/new", controllers.New)

	http.Handle("/", r)
}