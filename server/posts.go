package server

import (
	"blog/proto"
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (b *BlogService) ListPosts(ctx context.Context, req *proto.ListPostsRequest) (*proto.ListPostsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPosts not implemented")
}
func (b *BlogService) GetPost(ctx context.Context, req *proto.GetPostRequest) (*proto.Post, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPost not implemented")
}
func (b *BlogService) CreatePost(ctx context.Context, req *proto.CreatePostRequest) (*proto.Post, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePost not implemented")
}
func (b *BlogService) UpdatePost(ctx context.Context, req *proto.UpdatePostRequest) (*proto.Post, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePost not implemented")
}
func (b *BlogService) DeletePost(ctx context.Context, req *proto.DeletePostRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePost not implemented")
}
