package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/mnabil1718/go-database-module/entity"
)

type CommentRepositoryImpl struct {
	DB *sql.DB
}

func NewCommentRepository(db *sql.DB) CommentRepository {
	return &CommentRepositoryImpl{
		DB: db,
	}
}

func (repository *CommentRepositoryImpl) Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	query := "INSERT INTO comments (email, comment) VALUES ($1, $2) RETURNING id"
	var lastInsertedId int64
	err := repository.DB.QueryRowContext(ctx, query, comment.Email, comment.Comment).Scan(&lastInsertedId)
	if err != nil {
		return comment, err
	}

	comment.Id = lastInsertedId
	return comment, nil
}

func (repository *CommentRepositoryImpl) FindById(ctx context.Context, id int64) (entity.Comment, error) {
	query := "SELECT id, email, comment FROM comments WHERE id = $1"
	comment := entity.Comment{}
	err := repository.DB.QueryRowContext(ctx, query, id).Scan(&comment.Id, &comment.Email, &comment.Comment)
	if err != nil {
		return comment, errors.New("comment not found")
	}

	return comment, nil
}

func (repository *CommentRepositoryImpl) FindAll(ctx context.Context) ([]entity.Comment, error) {
	query := "SELECT id, email, comment FROM comments"
	var comments []entity.Comment
	rows, err := repository.DB.QueryContext(ctx, query)
	if err != nil {
		return comments, err
	}
	defer rows.Close()

	for rows.Next() {
		var comment entity.Comment
		err = rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		if err != nil {
			return comments, err
		}
		comments = append(comments, comment)
	}

	return comments, nil
}
