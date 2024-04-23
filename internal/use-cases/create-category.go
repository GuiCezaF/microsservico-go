package use_cases

import (
	"log"

	"github.com/GuiCezaF/microsservico-go/internal/entities"
)

type createCategoryUseCase struct {
	// db
}

func NewCreateCategoryUseCase() *createCategoryUseCase {
	return &createCategoryUseCase{}
}

func (u *createCategoryUseCase) Execute(name string) error {
	category, err := entities.NewCategory(name)

	if err != nil {
		return err
	}

	// TODO: Persist entity to db

	log.Println(category)

	return nil
}
