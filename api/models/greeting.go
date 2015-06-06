package models

import (
	"time"
	"github.com/drborges/datastore-model"
)

type Greeting struct {
	db.Model 	      `db:"Greetings"`
	Author  string
	Content string    `json:"content" db:id`
	Date    time.Time
}

type Greetings []*Greeting

func (this Greetings) ByAuthor(author string) *db.Query {
	return db.QueryFor(new(Greeting)).Filter("Author=", author)
}