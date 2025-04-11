package repository

import (
	"context"
	"database/sql"
	"sync"
)

type PostRepository struct {
	db    *sql.DB
	mutex sync.Mutex
}

func NewPostRepository(db *sql.DB) *PostRepository {
	return &PostRepository{
		db: db,
	}
}

type PostRepositoryCreateParams struct {
	Title   string
	Content string
}

func (r *PostRepository) Create(ctx context.Context, params PostRepositoryCreateParams) (int, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	query := `INSERT INTO posts (title, content) VALUES ($1, $2) RETURNING id`
	var id int32
	err := r.db.QueryRowContext(ctx, query, params.Title, params.Content).Scan(&id)
	if err != nil {
		return 0, err
	}

	return int(id), nil
}
