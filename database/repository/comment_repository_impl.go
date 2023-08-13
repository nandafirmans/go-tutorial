package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/nandafirmans/go-tutorial/database/entity"
)

type commentRepositoryImpl struct {
	DB *sql.DB
}

func NewCommentRepository(db *sql.DB) CommentRepository {
	newCommentRepo := commentRepositoryImpl{DB: db}
	return &newCommentRepo
}

func (repo *commentRepositoryImpl) Insert(
	ctx context.Context,
	comment entity.Comment,
) (entity.Comment, error) {
	script := "INSERT INTO comments(email, comment) VALUES(?, ?)"

	result, err := repo.DB.ExecContext(ctx, script, comment.Email, comment.Comment)
	if err != nil {
		return comment, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return comment, err
	}

	comment.Id = int32(id)

	return comment, nil
}

func (repo *commentRepositoryImpl) FindById(
	ctx context.Context,
	id int32,
) (entity.Comment, error) {
	script := "SELECT id, email, comment FROM comments WHERE id = ? LIMIT 1"

	rows, err := repo.DB.QueryContext(ctx, script, id)
	comment := entity.Comment{}

	if err != nil {
		return comment, err
	}
	defer rows.Close()

	if rows.Next() {
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		return comment, nil
	} else {
		return comment, errors.New("Id " + string(id) + " not found")
	}
}

func (repo *commentRepositoryImpl) FindAll(ctx context.Context) ([]entity.Comment, error) {
	script := "SELECT id, email, comment FROM comments"

	rows, err := repo.DB.QueryContext(ctx, script)
	comments := []entity.Comment{}

	if err != nil {
		return comments, err
	}
	defer rows.Close()

	for rows.Next() {
		comment := entity.Comment{}
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		comments = append(comments, comment)

	}

	return comments, nil
}
