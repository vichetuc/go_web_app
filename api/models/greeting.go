package models

import (
	"time"
	"github.com/drborges/datastore-model"
)

type Greeting struct {
	db.Model 	      `db:"Greetings"`
	Author  string	  `json:"author"`
	Content string    `json:"content" db:id`
	Date    time.Time
}

type Greetings []*Greeting

func (this Greetings) ByAuthor(author string) *db.Query {
	return db.From(new(Greeting)).Filter("Author=", author)
}

func (this Greetings) GetAll() *db.Query {
	return db.From(new(Greeting))
}