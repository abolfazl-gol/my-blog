package server

import (
	"blog/models"
	"blog/pkg/invalid"
	"blog/proto"
	"context"
	"log"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (b *BlogService) ListPosts(ctx context.Context, req *proto.ListPostsRequest) (*proto.ListPostsResponse, error) {
	user := ctx.Value("user").(*models.User)

	blog, err := models.GetBlog(&models.GetBlogOption{UserID: user.ID, ID: req.BlogId})
	if err != nil {
		log.Println("[ListPosts] can't GetBlog:", err)
		if invalid.IsErrorNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "%v", err)
		}
		return nil, status.Errorf(codes.Internal, "%v", err)
	}

	opts := &models.FindOptions{
		BlogID:  blog.ID,
		Page:    int(req.Page),
		OrderBy: req.OrderBy,
		SortBy:  req.SortBy,
	}

	posts, err := models.FindPosts(opts)
	if err != nil {
		log.Println("[ListPosts] can't FindPosts:", err)
		if invalid.IsErrorNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "%v", err)
		}
		return nil, status.Errorf(codes.Internal, "%v", err)

	}

	protoPost := make([]*proto.Post, 0)
	for _, post := range posts {
		protoPost = append(protoPost, post.ToProto())
	}

	nextPage := req.Page + 1
	if len(posts) < opts.Limit {
		nextPage = 0
	}

	return &proto.ListPostsResponse{Posts: protoPost, NextPage: nextPage}, nil
}

func (b *BlogService) GetPost(ctx context.Context, req *proto.GetPostRequest) (*proto.Post, error) {
	user := ctx.Value("user").(*models.User)

	blog, err := models.GetBlog(&models.GetBlogOption{UserID: user.ID, ID: req.BlogId})
	if err != nil {
		log.Println("[GetPost] can't GetBlog:", err)
		if invalid.IsErrorNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "%v", err)
		}
		return nil, status.Errorf(codes.Internal, "%v", err)
	}

	post, err := models.GetPost(&models.GetPostOption{ID: req.PostId, BlogID: blog.ID})
	if err != nil {
		log.Println("[GetPost] can't GetPost:", err)
		if invalid.IsErrorNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "%v", err)
		}
		return nil, status.Errorf(codes.Internal, "%v", err)
	}

	return post.ToProto(), nil
}

func (b *BlogService) CreatePost(ctx context.Context, req *proto.CreatePostRequest) (*proto.Post, error) {
	// user := ctx.Value("user").(*models.User)

	return nil, status.Errorf(codes.Unimplemented, "method CreatePost not implemented")
}

func (b *BlogService) UpdatePost(ctx context.Context, req *proto.UpdatePostRequest) (*proto.Post, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePost not implemented")
}

func (b *BlogService) DeletePost(ctx context.Context, req *proto.DeletePostRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePost not implemented")
}
