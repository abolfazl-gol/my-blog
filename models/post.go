package models

import (
	"blog/pkg/invalid"
	"blog/proto"
	"strings"
	"time"
)

type Post struct {
	ID        int64
	BlogID    int64
	Title     string
	Body      string
	Likes     int32
	CreatedAt time.Time `xorm:"<-"`
	UpdatedAt time.Time `xorm:"<-"`
}

type GetPostOption struct {
	ID     int64
	BlogID int64
}

func (p *Post) ToProto() *proto.Post {
	return &proto.Post{
		Id:     p.ID,
		BlogId: p.BlogID,
		Title:  p.Title,
		Body:   p.Body,
		Likes:  p.Likes,
	}
}

func (p *Post) TableName() string { return "posts" }

func FindPosts(opts *FindOptions) ([]*Post, error) {
	opts.Parse()

	sess := engine.NewSession()
	defer sess.Close()

	if opts.BlogID != 0 {
		sess.Where("blog_id = ?", opts.BlogID)
	}

	switch strings.ToLower(opts.OrderBy) {
	case "asc":
		sess.Asc(opts.SortBy)
	case "desc":
		sess.Desc(opts.SortBy)
	default:
		sess.Desc(opts.SortBy)

	}

	sess.Limit(opts.Limit, opts.Offset)

	posts := make([]*Post, 0)
	if err := sess.Find(&posts); err != nil {
		return nil, err
	}

	return posts, nil
}

func GetPost(opts *GetPostOption) (*Post, error) {
	return getPost(engine, opts)
}

func getPost(e Engine, opts *GetPostOption) (*Post, error) {
	sess := engine.NewSession()
	defer sess.Close()

	if opts.ID != 0 {
		sess.ID(opts.ID)
	}

	if opts.BlogID != 0 {
		sess.Where("blog_id = ?", opts.BlogID)
	}

	post := new(Post)
	has, err := sess.Get(post)

	if err != nil {
		return nil, err
	}

	if !has {
		return nil, invalid.ErrNotFound{"Post", "id", opts.ID}
	}

	return post, nil
}
