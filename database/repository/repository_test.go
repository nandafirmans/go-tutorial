package repository

import (
	"context"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"

	"github.com/nandafirmans/go-tutorial/database"
	"github.com/nandafirmans/go-tutorial/database/entity"
)

func TestCommentInsert(t *testing.T) {
	db := database.GetConnection()
	ctx := context.Background()
	commentRepository := NewCommentRepository(db)

	newComment := entity.Comment{
		Email:   "afif@email.com",
		Comment: "pidapidapidapidapidapida",
	}

	result, err := commentRepository.Insert(ctx, newComment)
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully Insert:", result)
}

func TestFindById(t *testing.T) {
	db := database.GetConnection()
	ctx := context.Background()
	commentRepository := NewCommentRepository(db)

	comment, err := commentRepository.FindById(ctx, 32)
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully Find By Id:", comment)
}

func TestFindAll(t *testing.T) {
	db := database.GetConnection()
	ctx := context.Background()
	commentRepository := NewCommentRepository(db)

	comments, err := commentRepository.FindAll(ctx)
	if err != nil {
		panic(err)
	}

	for _, comment := range comments {
		fmt.Println(comment)
	}
}
