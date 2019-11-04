package server

import (
	"blog/models"
	"blog/pkg/invalid"
	"blog/proto"
	"context"
	"log"
	"strings"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/gosimple/slug"
	"github.com/rezam90/goutils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (b *BlogService) ListBlogs(ctx context.Context, req *proto.ListBlogsRequest) (*proto.ListBlogsResponse, error) {
	user := ctx.Value("user").(*models.User)

	options := &models.FindBlogOptions{
		UserID:  user.ID,
		Page:    int(req.Page),
		OrderBy: req.OrderBy,
		SortBy:  req.SortBy,
	}

	blogs, err := models.FindBlogs(options)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", err)
	}

	protoBlogs := make([]*proto.Blog, 0)
	for _, blogs := range blogs {
		protoBlogs = append(protoBlogs, blogs.ToProto())
	}

	nextPage := req.Page + 1
	if len(blogs) < options.Limit {
		nextPage = 0
	}

	return &proto.ListBlogsResponse{Blogs: protoBlogs, NextPage: nextPage}, nil
}
func (b *BlogService) GetBlog(ctx context.Context, req *proto.GetBlogRequest) (*proto.Blog, error) {
	user := ctx.Value("user").(*models.User)
	blog, err := models.GetBlog(&models.GetBlogOption{UserID: user.ID, ID: req.BlogId})

	if err != nil {
		if invalid.IsErrorNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "%v", err)
		}
		return nil, status.Errorf(codes.Internal, "%v", err)

	}
	return blog.ToProto(), nil
}

func (b *BlogService) CreateBlog(ctx context.Context, req *proto.CreateBlogRequest) (*proto.Blog, error) {
	user := ctx.Value("user").(*models.User)

	blog := &models.Blog{
		Name:        req.Name,
		Title:       req.Title,
		Description: req.Description,
		Slug:        slug.Make(req.Name),
		UserID:      user.ID,
	}

	if err := models.CreateBlog(blog); err != nil {
		if invalid.IsErrDuplicate(err) {
			return nil, status.Errorf(codes.InvalidArgument, "can't create blog: %v", err)
		}
		return nil, status.Errorf(codes.Internal, "%v", err)

	}

	return blog.ToProto(), nil
}
func (b *BlogService) UpdateBlog(ctx context.Context, req *proto.UpdateBlogRequest) (*proto.Blog, error) {
	user := ctx.Value("user").(*models.User)

	if req.Blog == nil {
		return nil, status.Error(codes.InvalidArgument, "Account must be set")
	}

	if len(req.UpdateMask.Paths) == 0 {
		return nil, status.Error(codes.InvalidArgument, "update_mask is empty, should be like ['blog.status']")
	}

	blog, err := models.GetBlog(&models.GetBlogOption{ID: req.Blog.Id, UserID: user.ID})
	if err != nil {
		log.Println("[UpdateBlog] can't GetBlog:", err)
		if invalid.IsErrorNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "%v", err)
		}
		return nil, status.Errorf(codes.Internal, "can't get blog: %v", err)
	}

	forbiddens := []string{"id", "user_id"}
	colums := make([]string, 0)
	for _, path := range req.UpdateMask.GetPaths() {
		if strings.HasPrefix(path, "blog.") {
			colum := path[strings.Index(path, ".")+1:]
			if goutils.SliceExists(forbiddens, colum) {
				continue
			}
			colums = append(colums, colum)
		}
	}

	if req.Blog.Name != "" && req.Blog.Name != blog.Name {
		blog.Slug = slug.Make(req.Blog.Name)
		colums = append(colums, "slug")
	}

	blog.Name = req.Blog.Name
	blog.Title = req.Blog.Title
	blog.Description = req.Blog.Description

	if err := models.UpdateBlog(blog, colums...); err != nil {
		log.Println("[UpdateBlog] can't UpdateBlog:", err)
		if invalid.IsErrDuplicate(err) {
			return nil, status.Errorf(codes.InvalidArgument, "%v", err)
		}
		return nil, status.Errorf(codes.Internal, "can't UpdateBlog: %v", err)
	}

	newBlog, err := models.GetBlog(&models.GetBlogOption{ID: req.Blog.Id})
	if err != nil {
		log.Println("[UpdateBlog] can't GetBlog:", err)
		if invalid.IsErrorNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "%v", err)
		}
		return nil, status.Errorf(codes.Internal, "can't get blog: %v", err)
	}

	return newBlog.ToProto(), nil
}
func (b *BlogService) DeleteBlog(ctx context.Context, req *proto.DeleteBlogRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBlog not implemented")
}
