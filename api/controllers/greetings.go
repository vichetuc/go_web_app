package controllers
import (
	"github.com/gin-gonic/gin"
	"github.com/juliofarah/go_web_app/api/models"
	"net/http"
	"fmt"
)

type Greetings struct {
	Controller
}

func (this *Greetings) List(c *gin.Context) {
	author := c.Query("author")
	greetings := models.Greetings{}

	if err := this.Datasource.Query(greetings.ByAuthor(author)).All(&greetings); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H {
			"message" : fmt.Sprintf("%v", err),
		})
		return
	}

	c.JSON(http.StatusOK, greetings)
}