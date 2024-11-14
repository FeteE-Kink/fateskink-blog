package models

import "time"

type Tag struct {
	Id        int32
	Name      string
	UserId    int32
	Articles  []*Article `gorm:"many2many:articles_tags;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
