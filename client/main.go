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

	blog, err := client.CreateBlog(ctx, &proto.CreateBlogRequest{Name: "MyBlo//gs", Title: "blog ", Description: "Hi! Welcom to blog"})
	// blogs, err := client.ListBlogs(ctx, &proto.ListBlogsRequest{OrderBy: "asc"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(blog)
	// blog, err := client.GetBlog(ctx, &proto.GetBlogRequest{BlogId: 11})
	// blog, err := client.UpdateBlog(ctx, &proto.UpdateBlogRequest{
	// 	Blog: &proto.Blog{
	// 		Id:          11,
	// 		Name:        "Good by",
	// 		Title:       "welcom",
	// 		Description: "heeeellllooo",
	// 	},
	// 	UpdateMask: &field_mask.FieldMask{
	// 		Paths: []string{"blog.name", "blog.title", "blog.description"},
	// 	},
	// })
	// blog, err := client.DeleteBlog(ctx, &proto.DeleteBlogRequest{BlogId: 14})
	// posts, err := client.ListPosts(ctx, &proto.ListPostsRequest{OrderBy: "asc", BlogId: 13})
	// post, err := client.GetPost(ctx, &proto.GetPostRequest{PostId: 9, BlogId: 6})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(post)

}
