package repository

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/mnabil1718/go-database-module/database"
	"github.com/mnabil1718/go-database-module/entity"
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
