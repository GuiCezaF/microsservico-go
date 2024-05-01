package repositories

import "github.com/GuiCezaF/microsservico-go/internal/entities"

type ICategoryRepository interface {
	Save(category *entities.Category) error
	List() ([]*entities.Category, error)
}
