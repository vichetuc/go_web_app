package guestbook

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/juliofarah/go_web_app/api/controllers"
)

func init() {

	r := gin.Default()

	r.GET("/greetings", controllers.AllGreetings)
	r.POST("/greetings", controllers.New)

	r.Static("/", "../static")

	http.Handle("/", r)
}