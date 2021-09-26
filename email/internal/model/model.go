package model

import "time"

type Email struct {
	Id          string
	Email       string
	To          string
	From        string
	Body        string
	Subject     string
	ContentType string
	CreatedAt   time.Time
}

