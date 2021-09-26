package model

import "time"

type Posts struct {
	Id int
	AuthorId int
	Title string
	MetaTitle string
	Slug string
	Summary string
	Published int
	CreatedAt time.Time
	UpdatedAt time.Time
	PublishedAt time.Time
	Content string
}