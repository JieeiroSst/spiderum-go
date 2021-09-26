package model

import "time"

type Ip struct {
	Id     int `gorm:"primary_key" json:"id"`
	Ip     string
	Method string
	RequestAt time.Time `json:"created_at"`
}