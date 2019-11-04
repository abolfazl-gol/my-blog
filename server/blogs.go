package server

import (
	"blog/models"
	"blog/pkg/invalid"
	"blog/proto"
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (b *BlogService) ListBlogs(ctx context.Context, req *proto.ListBlogsRequest) (*proto.ListBlogsResponse, error) {
	// user := ctx.Value("user_id").(*models.User)
	return nil, status.Errorf(codes.Unimplemented, "method ListBlogs not implemented")
}
func (b *BlogService) GetBlog(ctx context.Context, req *proto.GetBlogRequest) (*proto.Blog, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlog not implemented")
}
func (b *BlogService) CreateBlog(ctx context.Context, req *proto.CreateBlogRequest) (*proto.Blog, error) {
	user := ctx.Value("user").(*models.User)

	blog := &models.Blog{
		Name:        req.Name,
		Title:       req.Title,
		Description: req.Description,
		Slug:        req.Name,
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
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBlog not implemented")
}
func (b *BlogService) DeleteBlog(ctx context.Context, req *proto.DeleteBlogRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBlog not implemented")
}
