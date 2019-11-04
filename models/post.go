package models

import "time"

type Post struct {
	ID        int64
	Title     string
	Body      string
	BlogID    int64
	Likes     int
	CreatedAt time.Time `xorm:"<-"`
	UpdatedAt time.Time `xorm:"<-"`
}

func (p *Post) TableName() string { return "posts" }
