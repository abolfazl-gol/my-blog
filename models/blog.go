package models

import (
	"blog/pkg/invalid"
	"blog/proto"
	"strings"
	"time"
)

type Blog struct {
	ID          int64
	Name        string
	Title       string
	Description string
	UserID      int64
	Slug        string
	CreatedAt   time.Time `xorm:"<-"`
	UpdatedAt   time.Time `xorm:"<-"`
}

func (b *Blog) ToProto() *proto.Blog {
	return &proto.Blog{
		Id:          b.ID,
		Name:        b.Name,
		Title:       b.Title,
		Description: b.Description,
		UserId:      b.UserID,
		Slug:        b.Slug,
	}
}

func (b *Blog) TableName() string { return "blogs" }

func CreateBlog(blog *Blog) error {
	return createBlog(engine, blog)
}

func createBlog(e Engine, blog *Blog) error {
	if _, err := e.Insert(blog); err != nil {
		if strings.Contains(err.Error(), "1062") {
			return invalid.ErrDuplicate{"Blog", "name", blog.Name}
		}
		return err
	}

	return nil
}
