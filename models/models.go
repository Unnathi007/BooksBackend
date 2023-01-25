package models

import (
	"time"
)

type Book struct {
	Id     int       `json:"id"`
	Name   string    `json:"name"`
	Author string    `json:"author"`
	Pages  int       `json:"pages"`
	DOP    time.Time `json:"publication_date"`
}
