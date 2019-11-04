package main

import (
	"blog/proto"
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	client := proto.NewApiClient(conn)
	ctx := context.Background()
	ctx = metadata.AppendToOutgoingContext(ctx, "token", "c19a4e97-97ef-47b1-a46d-cad0a76c9d75")
	// user, err := client.Register(ctx, &proto.RegisterRequest{Name: "reza", Email: "rezam.golm@gmail.ir", Password: "123456"})
	// // user, err := client.Login(ctx, &proto.LoginRequest{Email: "alireza.golm@gmail.ir", Password: "123456"})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(user)

	blog, err := client.CreateBlog(ctx, &proto.CreateBlogRequest{Name: "بلاگ", Title: "بلاگر ", Description: "Hi! Welcom to blog"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(blog)
}
