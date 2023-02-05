package repository

import "github.com/nandafirmans/go-tutorial/unit-test/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
