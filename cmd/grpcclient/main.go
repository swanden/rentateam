package main

import (
	"context"
	"github.com/swanden/rentateam/api/grpcpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := grpcpb.NewPostsClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Create(ctx, &grpcpb.CreateRequest{
		Post: &grpcpb.Post{
			Title:     "Post title 16",
			Body:      "Post body",
			Tags:      []string{"First", "Second"},
			CreatedAt: timestamppb.Now(),
		},
	})
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	log.Printf("Post ID: %d", r.GetId())

	ctxAll, cancelAll := context.WithTimeout(context.Background(), time.Second)
	defer cancelAll()
	ar, err := c.All(ctxAll, &grpcpb.AllRequest{})
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	log.Printf("Posts: %v", ar.Posts)
}
