package guestbook

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/juliofarah/go_web_app/api/controllers"
)

func init() {

	r := gin.New()

	r.GET("/", controllers.AllGreetings)
	r.GET("/json", controllers.GreetingsToJson)
	r.POST("/sign", controllers.Sign)

	http.Handle("/", r)
}