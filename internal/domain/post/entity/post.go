package entity

import (
	"github.com/jackc/pgtype"
)

type Post struct {
	ID        int
	Title     string
	Body      string
	Tags      []string
	CreatedAt pgtype.Timestamp
}

func New(title, body string, tags []string, createdAt pgtype.Timestamp) Post {
	return Post{
		Title:     title,
		Body:      body,
		Tags:      tags,
		CreatedAt: createdAt,
	}
}
