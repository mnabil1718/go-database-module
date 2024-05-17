package repository

import (
	"context"

	"github.com/mnabil1718/go-database-module/entity"
)

type CommentRepository interface {
	Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error)
	FindById(ctx context.Context, id int64) (entity.Comment, error)
	FindAll(ctx context.Context) ([]entity.Comment, error)
}
