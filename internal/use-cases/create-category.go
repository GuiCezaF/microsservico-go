package use_cases

import (
	"log"

	"github.com/GuiCezaF/microsservico-go/internal/entities"
	"github.com/GuiCezaF/microsservico-go/internal/repositories"
)

type createCategoryUseCase struct {
	repository repositories.ICategoryRepository
}

func NewCreateCategoryUseCase(repository repositories.ICategoryRepository) *createCategoryUseCase {
	return &createCategoryUseCase{repository}
}

func (u *createCategoryUseCase) Execute(name string) error {
	category, err := entities.NewCategory(name)

	if err != nil {
		return err
	}

	log.Println(category)
	err = u.repository.Save(category)

	if err != nil {
		return err
	}

	return nil
}
