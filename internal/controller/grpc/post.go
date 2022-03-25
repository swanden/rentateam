package grpc

import (
	"context"
	"fmt"
	"github.com/jackc/pgtype"
	"github.com/pkg/errors"
	"github.com/swanden/rentateam/api/grpcpb"
	"github.com/swanden/rentateam/internal/domain"
	"github.com/swanden/rentateam/internal/domain/post/repository"
	"github.com/swanden/rentateam/internal/domain/post/usecase"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"os"
	"strings"
)

type PostController struct {
	grpcpb.UnimplementedPostsServer
	repo    repository.PostRepository
	useCase *usecase.PostUseCase
}

func NewPost(repo repository.PostRepository, useCase *usecase.PostUseCase) *PostController {
	return &PostController{
		repo:    repo,
		useCase: useCase,
	}
}

var _ grpcpb.PostsServer = (*PostController)(nil)

func (p *PostController) Create(ctx context.Context, request *grpcpb.CreateRequest) (*grpcpb.CreateResponse, error) {
	if strings.TrimSpace(request.Post.Title) == "" {
		return nil, status.Error(codes.InvalidArgument, "Post title is empty")
	}
	if strings.TrimSpace(request.Post.Body) == "" {
		return nil, status.Error(codes.InvalidArgument, "Post body is empty")
	}

	var createdAt pgtype.Timestamp
	err := createdAt.Set(request.Post.CreatedAt.AsTime())
	if err != nil {
		fmt.Fprintf(os.Stderr, "grpc - v1 - postController - create - Set: %v\n", err)
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	createDTO := usecase.CreateDTO{
		Title:     request.Post.Title,
		Body:      request.Post.Body,
		Tags:      request.Post.Tags,
		CreatedAt: createdAt,
	}

	id, err := p.useCase.Create(ctx, createDTO)

	if errors.Cause(err) == domain.Error {
		return nil, status.Error(codes.AlreadyExists, err.Error())
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "grpc - v1 - postController - Create - Cause: %v\n", err)
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	return &grpcpb.CreateResponse{Id: int32(id)}, nil
}

func (p *PostController) All(ctx context.Context, request *grpcpb.AllRequest) (*grpcpb.AllResponse, error) {
	posts, err := p.repo.All(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "grpc - v1 - postController - All - All: %v\n", err)
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	var response grpcpb.AllResponse
	for _, post := range posts {
		respPost := grpcpb.AllPost{
			Id:        int32(post.ID),
			Title:     post.Title,
			Body:      post.Body,
			Tags:      post.Tags,
			CreatedAt: timestamppb.New(post.CreatedAt.Time),
		}
		response.Posts = append(response.Posts, &respPost)
	}

	return &response, nil
}
