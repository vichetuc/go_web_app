package controllers
import (
	"net/http"
	"appengine"
	"github.com/juliofarah/go_web_app/api/models"
	"time"
	"appengine/user"
	"github.com/drborges/datastore-model"
	"github.com/gin-gonic/gin"
)

func Sign(c *gin.Context) {

	request := c.Request
	context := appengine.NewContext(request)

	greeting := new(models.Greeting)
	greeting.Content = request.FormValue("content")
	greeting.Date = time.Now()

	if u := user.Current(context); u != nil {
		greeting.Author = u.String()
	}

	if err := db.NewDatastore(context).Create(greeting); err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
	c.Redirect(301, "/")
}

