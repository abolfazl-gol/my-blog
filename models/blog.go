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

type GetBlogOption struct {
	ID     int64
	UserID int64
}

func (b *Blog) ToProto() *proto.Blog {
	return &proto.Blog{
		Id:          b.ID,
		Name:        b.Name,
		Title:       b.Title,
		Description: b.Description,
		UserId:      b.UserID,
		// Slug:        b.Slug,
	}
}

func (b *Blog) TableName() string { return "blogs" }

func GetBlog(opt *GetBlogOption) (*Blog, error) {
	return getBlog(engine, opt)
}

func CreateBlog(blog *Blog) error {
	return createBlog(engine, blog)
}

func UpdateBlog(blog *Blog, cols ...string) error {
	return updateBlog(engine, blog, cols...)
}

func DeleteBlog(id int64) error {
	return deleteBlog(engine, id)
}

func FindBlogs(opts *FindOptions) ([]*Blog, error) {
	opts.Parse()

	sess := engine.NewSession()
	defer sess.Close()

	if opts.UserID != 0 {
		sess.Where("user_id = ?", opts.UserID)
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

	blogs := make([]*Blog, 0)
	if err := sess.Find(&blogs); err != nil {
		return nil, err
	}

	return blogs, nil
}

func getBlog(e Engine, opt *GetBlogOption) (*Blog, error) {
	sess := engine.NewSession()
	defer sess.Close()

	if opt.ID != 0 {
		sess.ID(opt.ID)
	}

	if opt.UserID != 0 {
		sess.Where("user_id = ?", opt.UserID)
	}

	blog := new(Blog)
	has, err := sess.Get(blog)
	if err != nil {
		return nil, err
	}

	if !has {
		return nil, invalid.ErrNotFound{"Blog", "id", opt.ID}
	}

	return blog, nil
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

func updateBlog(e Engine, blog *Blog, cols ...string) error {
	if _, err := e.ID(blog.ID).UseBool().Cols(cols...).Update(blog); err != nil {
		if strings.Contains(err.Error(), "1062") {
			return invalid.ErrDuplicate{"Blog", "name", blog.Name}
		}
		return err
	}
	return nil
}

func deleteBlog(e Engine, id int64) error {
	if _, err := e.Delete(&Blog{ID: id}); err != nil {
		return err
	}
	return nil
}
