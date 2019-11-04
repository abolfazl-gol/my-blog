package server

import (
	"blog/models"
	"blog/pkg/invalid"
	"blog/pkg/validate"
	"blog/proto"
	"context"

	"github.com/gofrs/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (b *BlogService) Register(ctx context.Context, req *proto.RegisterRequest) (*proto.User, error) {
	if err := validate.RegisterReq(req); err != nil {
		return nil, err
	}
	user := &models.User{
		Email: req.Email,
		Name:  req.Name,
		Token: uuid.Must(uuid.NewV4()).String(),
	}

	user.GenerateHash(req.Password)

	if err := models.CreateUser(user); err != nil {
		if invalid.IsErrDuplicate(err) {
			return nil, status.Errorf(codes.InvalidArgument, "can't create user: %v", err)
		}
		return nil, status.Errorf(codes.Internal, "%v", err)
	}

	return user.ToProto(), nil
}

func (b *BlogService) Login(ctx context.Context, req *proto.LoginRequest) (*proto.User, error) {
	user, err := models.GetUserByEmail(req.Email)
	if err != nil {
		if invalid.IsErrorNotFound(err) {
			return nil, status.Errorf(codes.NotFound, "%v", err)
		}
		return nil, status.Errorf(codes.Internal, "internal:%v", err)
	}
	if !user.CheckPasswordHash(req.Password) {
		return nil, status.Error(codes.Unauthenticated, " invalid email/password")
	}
	return user.ToProto(), nil
}
