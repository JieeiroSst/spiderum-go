package model

import "time"

type Profiles struct {
	Id 	int `gorm:"primary_key" json:"id"`
	UserId int `json:"user_id"`
	FirstName string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	LastName string `json:"last_name"`
	Mobile string `json:"mobile"`
	Email string `json:"email"`
	RegisteredAt time.Time `json:"registered_at"`
	Profile string `json:"profile"`
}

type Posts struct {
	Id int `gorm:"primary_key" json:"id"`
	AuthorId int `json:"author_id"`
	Title string `json:"title"`
	MetaTitle string `json:"meta_title"`
	Slug string `json:"slug"`
	Summary string `json:"summary"`
	Published int `json:"published"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:null"`
	PublishedAt time.Time `json:"published_at" gorm:"default:null"`
	Content string `json:"content"`
}

type PostMetas struct {
	Id int `gorm:"primary_key" json:"id"`
	PostId int `json:"post_id"`
	TextKey string `json:"text_key"`
	Content string `json:"content"`
}

type PostComments struct {
	Id int `gorm:"primary_key" json:"id"`
	PostId int `json:"post_id"`
	ParentId int `json:"parent_id"`
	Title string `json:"title"`
	Published int `json:"published"`
	CreatedAt time.Time `json:"created_at"`
	PublishedAt time.Time `json:"published_at" gorm:"default:null"`
	Content string `json:"content"`
}

type Categories struct {
	Id int `gorm:"primary_key" json:"id"`
	ParentId int `json:"parent_id"`
	Title string `json:"title"`
	MetaTitle string `json:"meta_title"`
	Slug string `json:"slug"`
	Content string `json:"content"`
}

