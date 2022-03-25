package repository

import (
	"context"
	"github.com/swanden/rentateam/internal/domain/post/entity"
)

type PostRepository interface {
	Save(ctx context.Context, post entity.Post) (int, error)
	All(ctx context.Context) ([]entity.Post, error)
	HasByTitle(ctx context.Context, title string) (bool, error)
}
