package controllers

import (
	"github.com/drborges/datastore-model"
	"github.com/gin-gonic/gin"
	"appengine"
	"github.com/juliofarah/go_web_app/api/services"
)

type Controller struct {
	GaeContext appengine.Context
	Datasource db.Datasource
}

func (this *Controller) Register(c *gin.Context) {
	this.GaeContext = services.Gae{c.Request}.NewContext()
	this.Datasource = db.NewDatastore(this.GaeContext)
}