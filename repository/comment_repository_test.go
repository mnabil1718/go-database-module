package repository

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/mnabil1718/go-database-module/database"
	"github.com/mnabil1718/go-database-module/entity"
	"github.com/stretchr/testify/assert"
)

func TestCommentInsert(t *testing.T) {
	db := database.GetConnection()
	defer db.Close()

	ctx := context.Background()
	commentRepository := NewCommentRepository(db)

	newComment := entity.Comment{
		Email:   "mnabil1718@gmail.com",
		Comment: sql.NullString{String: "FIRST", Valid: true},
	}

	comment, err := commentRepository.Insert(ctx, newComment)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(comment)
}

func TestCommentFindById(t *testing.T) {
	db := database.GetConnection()
	defer db.Close()

	ctx := context.Background()
	commentRepository := NewCommentRepository(db)

	comment, err := commentRepository.FindById(ctx, 16)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(comment)
}

func TestCommentFindAll(t *testing.T) {
	db := database.GetConnection()
	defer db.Close()

	ctx := context.Background()
	commentRepository := NewCommentRepository(db)

	comments, err := commentRepository.FindAll(ctx)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(comments)
}

func TestCommentFindByIdMock(t *testing.T) {
	CommentRepositoryImplMock := &CommentRepositoryImplMock{}
	var comment = entity.Comment{
		Id:      1,
		Email:   "joko@gmail.com",
		Comment: sql.NullString{String: "Komentar ke-0", Valid: true},
	}
	CommentRepositoryImplMock.Mock.On("FindById", context.Background(), int64(1)).Return(comment, nil)
	commentResult, err := CommentRepositoryImplMock.FindById(context.Background(), 1)

	assert.Nil(t, err)
	assert.NotNil(t, commentResult)

	assert.Equal(t, int64(1), commentResult.Id)
	assert.Equal(t, commentResult.Email, comment.Email)
	assert.Equal(t, commentResult.Comment, comment.Comment)
}

func TestCommentInsertMock(t *testing.T) {
	CommentRepositoryImplMock := &CommentRepositoryImplMock{}
	var comment = entity.Comment{
		Id:      1,
		Email:   "joko@gmail.com",
		Comment: sql.NullString{String: "Komentar ke-0", Valid: true},
	}
	CommentRepositoryImplMock.Mock.On("Insert", context.Background(), comment).Return(comment, nil)
	commentResult, err := CommentRepositoryImplMock.Insert(context.Background(), comment)

	assert.Nil(t, err)
	assert.NotNil(t, commentResult)

	assert.Equal(t, commentResult.Email, comment.Email)
	assert.Equal(t, commentResult.Comment, comment.Comment)
}

func TestCommentFindAllMock(t *testing.T) {
	CommentRepositoryImplMock := &CommentRepositoryImplMock{}
	comments := []entity.Comment{
		{
			Id:      1,
			Email:   "munafiq",
			Comment: sql.NullString{String: "Komentar ke-0", Valid: true},
		},
	}
	CommentRepositoryImplMock.Mock.On("FindAll", context.Background()).Return(comments, nil)
	commentsResult, err := CommentRepositoryImplMock.FindAll(context.Background())

	assert.Nil(t, err)
	assert.NotNil(t, commentsResult)
}
