package middleware

import (
	"blog/models"
	"blog/pkg/invalid"
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func Interceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	if info.FullMethod == "/proto.Api/Register" || info.FullMethod == "/proto.Api/Login" {
		return handler(ctx, req)
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.InvalidArgument, "metadata not set")
	}

	token, ok := md["token"]
	if !ok || len(token) == 0 {
		return nil, status.Error(codes.Unauthenticated, "token not set")
	}

	user, err := models.GetUserByToken(token[0])
	if err != nil {
		if invalid.IsErrorNotFound(err) {
			return nil, status.Errorf(codes.Unauthenticated, "%v", err)
		}

		return nil, status.Errorf(codes.Internal, "internal:%v", err)
	}

	ctx = context.WithValue(ctx, "user_id", user.ID)
	ctx = context.WithValue(ctx, "user", user)

	return handler(ctx, req)
}
