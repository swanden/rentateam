package usecase

import (
	"context"
	"github.com/jackc/pgtype"
	"github.com/pkg/errors"
	"github.com/swanden/rentateam/internal/domain"
	"github.com/swanden/rentateam/internal/domain/post/entity"
	"github.com/swanden/rentateam/internal/domain/post/repository"
)

var AlreadyExistsPost = errors.Wrap(domain.Error, "post with such title is already exists")

type PostUseCase struct {
	postRepository repository.PostRepository
}

func New(repository repository.PostRepository) *PostUseCase {
	return &PostUseCase{
		postRepository: repository,
	}
}

type CreateDTO struct {
	Title     string
	Body      string
	Tags      []string
	CreatedAt pgtype.Timestamp
}

func (p *PostUseCase) Create(ctx context.Context, dto CreateDTO) (int, error) {
	alreadyExists, err := p.postRepository.HasByTitle(ctx, dto.Title)
	if err != nil {
		return 0, err
	}
	if alreadyExists {
		return 0, AlreadyExistsPost
	}

	post := entity.New(dto.Title, dto.Body, dto.Tags, dto.CreatedAt)

	return p.postRepository.Save(ctx, post)
}
