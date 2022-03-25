package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/swanden/rentateam/internal/domain/post/entity"
	"github.com/swanden/rentateam/internal/domain/post/repository"
	"github.com/swanden/rentateam/pkg/postgres"
)

type PostRepositoryPostgres struct {
	*postgres.Postgres
}

func New(db *postgres.Postgres) *PostRepositoryPostgres {
	return &PostRepositoryPostgres{db}
}

var _ repository.PostRepository = (*PostRepositoryPostgres)(nil)

func (r *PostRepositoryPostgres) All(ctx context.Context) ([]entity.Post, error) {
	rows, err := r.Pool.Query(ctx, "SELECT * FROM posts")
	if err != nil {
		return nil, fmt.Errorf("PostRepository - All - r.Pool.Query: %w\n", err)
	}
	defer rows.Close()

	posts := make([]entity.Post, 0)

	for rows.Next() {
		p := entity.Post{}

		err = rows.Scan(&p.ID, &p.Title, &p.Body, &p.Tags, &p.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("PostRepository - All - rows.Scan: %w", err)
		}

		posts = append(posts, p)
	}

	return posts, nil
}

func (r *PostRepositoryPostgres) Save(ctx context.Context, post entity.Post) (int, error) {
	var id int
	err := r.Pool.QueryRow(ctx, "INSERT INTO posts(title, body, tags, created_at) VALUES($1, $2, $3, $4) RETURNING id", post.Title, post.Body, post.Tags, post.CreatedAt).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("PostRepositoryPostgres - Save - r.Pool.Exec: %w", err)
	}

	return id, nil
}

func (r *PostRepositoryPostgres) HasByTitle(ctx context.Context, title string) (bool, error) {
	var count int
	err := r.Pool.QueryRow(ctx, "SELECT count(*) FROM posts WHERE title=$1", title).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("PostRepository - HasByTitle - r.Pool.Query: %w\n", err)
	}

	return count > 0, nil
}

func NullString(s string) sql.NullString {
	if len(s) == 0 {
		return sql.NullString{}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}
