package repository

import (
	"context"
	"database/sql"
	"forum/internal/model"
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

func (r *PostRepository) List(ctx context.Context, limit, offset int) ([]*model.Post, error) {
	query := `SELECT id, title, content, created_at, updated_at FROM post ORDER BY created_at DESC LIMIT $1 OFFSET $2`
	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}

	var posts []*model.Post
	for rows.Next() {
		var post model.Post
		err := rows.Scan(&post.Id, &post.Title, &post.Content, &post.CreatedAt, &post.UpdatedAt)
		if err != nil {
			return nil, err
		}
		posts = append(posts, &post)
	}

	rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

func (r *PostRepository) Get(ctx context.Context, id int) (*model.Post, error) {
	query := `SELECT id, title, content, created_at, updated_at FROM post WHERE id = $1`
	var post model.Post
	err := r.db.QueryRowContext(ctx, query, id).Scan(&post.Id, &post.Title, &post.Content, &post.CreatedAt, &post.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &post, nil
}

type PostRepositoryCreateParams struct {
	Title   string
	Content string
}

func (r *PostRepository) Create(ctx context.Context, params PostRepositoryCreateParams) (int, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	query := `INSERT INTO post (title, content) VALUES ($1, $2) RETURNING id`
	var id int32
	err := r.db.QueryRowContext(ctx, query, params.Title, params.Content).Scan(&id)
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

type PostRepositoryUpdateParams struct {
	Id      int
	Title   string
	Content string
}

func (r *PostRepository) Update(ctx context.Context, params PostRepositoryUpdateParams) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	query := `UPDATE post SET title = $1, content = $2 WHERE id = $3`
	_, err := r.db.ExecContext(ctx, query, params.Title, params.Content, params.Id)

	return err
}

func (r *PostRepository) Delete(ctx context.Context, id int) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	query := `DELETE FROM post WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)

	return err
}
