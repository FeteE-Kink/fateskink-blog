package models

import "time"

type ArticleTag struct {
	Id        int32
	ArticleId int32
	TagId     int32
	CreatedAt time.Time
	UpdatedAt time.Time
}
