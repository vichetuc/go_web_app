package models

import (
	"time"
	"github.com/drborges/datastore-model"
)

type Greeting struct {
	db.Model 	      `db:"Greetings"`
	Author  string	  `json:"author" binding:"required"`
	Content string    `json:"content" binding:"required"`
	Date    time.Time
}

type Greetings []*Greeting

func (this Greetings) ByAuthor(author string) *db.Query {
	return db.From(new(Greeting)).Filter("Author=", author)
}

func (this Greetings) GetAll() *db.Query {
	return db.From(new(Greeting))
}

func (this Greetings) New(content string, author string) *Greeting {
	greeting := new(Greeting)
	greeting.Content = content
	greeting.Date = time.Now()
	greeting.Author = author

	return greeting
}

