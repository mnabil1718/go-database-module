package repository

import (
	"context"

	"github.com/mnabil1718/go-database-module/entity"
	"github.com/stretchr/testify/mock"
)

type CommentRepositoryImplMock struct {
	Mock mock.Mock
}

func (repository *CommentRepositoryImplMock) Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	arguments := repository.Mock.Called(ctx, comment)
	return arguments.Get(0).(entity.Comment), arguments.Error(1)
}

func (repository *CommentRepositoryImplMock) FindById(ctx context.Context, id int64) (entity.Comment, error) {
	arguments := repository.Mock.Called(ctx, id)
	return arguments.Get(0).(entity.Comment), arguments.Error(1)
}

func (repository *CommentRepositoryImplMock) FindAll(ctx context.Context) ([]entity.Comment, error) {
	arguments := repository.Mock.Called(ctx)
	return arguments.Get(0).([]entity.Comment), arguments.Error(1)
}
