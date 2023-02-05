package service

import (
	"errors"

	"github.com/nandafirmans/go-tutorial/unit-test/entity"
	"github.com/nandafirmans/go-tutorial/unit-test/repository"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (service CategoryService) Get(id string) (*entity.Category, error) {
	category := service.Repository.FindById(id)
	if category == nil {
		return category, errors.New("Category Not Found")
	}
	return category, nil
}
