package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id        string `gorm:"primary_key; column:id"`
	Email     string `gorm:"not null; unique; size:32"`
	Password  string `gorm:"not null; size:20"`
	Role      string `gorm:"not null; size:10"`
	CreatedAt string `gorm:"column:created_at"`
	UpdatedAt string `gorm:"column:updated_at"`
}

type CategoryArticle struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
	Url  string `json:"url"`
}
