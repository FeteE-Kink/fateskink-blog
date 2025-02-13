package models

import "time"

type Article struct {
	Id        int32
	Title     string
	Content   string
	UserId    int32
	Status    int32
	Slug      string
	Tags      []*Tag `gorm:"many2many:articles_tags;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
