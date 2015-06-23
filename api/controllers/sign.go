package controllers
import (
	"net/http"
	"appengine"
	"github.com/juliofarah/go_web_app/api/models"
	"appengine/user"
	"github.com/drborges/datastore-model"
	"github.com/gin-gonic/gin"
)

func Sign(c *gin.Context) {

	request := c.Request
	context := appengine.NewContext(request)

	content := request.FormValue("content")

	greetings := models.Greetings{}

	user := user.Current(context)
	greeting := greetings.New(content, user.String())

	if err := db.NewDatastore(context).Create(greeting); err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
	c.Redirect(301, "/")
}

