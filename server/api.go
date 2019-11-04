package server

import (
	mw "blog/middleware"
	"blog/proto"
	"fmt"
	"log"
	"net"

	"github.com/rezam90/grpcutils"
)

type BlogService struct{}

func Started() error {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	server := grpcutils.NewServerBuilder().
		AddUnaryInterceptor(mw.Interceptor).
		Server()

	proto.RegisterApiServer(server, &BlogService{})
	fmt.Println("server started on port", ":50051")
	return server.Serve(lis)
}
