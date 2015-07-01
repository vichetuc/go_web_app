package controllers

import (
	"net/http"
	"appengine"
	"github.com/juliofarah/go_web_app/api/models"
	"github.com/drborges/datastore-model"
	"github.com/gin-gonic/gin"
	"github.com/juliofarah/go_web_app/api/modelAndView"
	"appengine/user"
)

func Create(c *gin.Context) {
	//can I extract context and datastore ?
	//datastore should live inside the model?
	greetingParams := modelAndView.GreetingAsForm{}
	c.Bind(&greetingParams)
	persist(c, greetingParams.Content)
	c.String(201, "")
}

func List(c *gin.Context) {

	context := appengine.NewContext(c.Request)
	greetings := models.Greetings{}

	if err := db.NewDatastore(context).Query(greetings.GetAll()).All(&greetings); err != nil {
		statusInternalServerError(c, err)
	} else {
		c.JSON(http.StatusOK, greetings)
	}
}

func statusInternalServerError(c *gin.Context, err error) {
	http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
}

func persist(c *gin.Context, content string) {
	context := appengine.NewContext(c.Request)
	greetings := models.Greetings{}

	author := "guest"
	if u := user.Current(context); u != nil {
		author = u.String()
	}
	greeting := greetings.New(author, content)

	if err := db.NewDatastore(context).Create(greeting); err != nil {
		statusInternalServerError(c, err)
		return
	}
}
