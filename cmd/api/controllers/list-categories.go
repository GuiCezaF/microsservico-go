package controllers

import (
	"net/http"

	"github.com/GuiCezaF/microsservico-go/internal/repositories"
	use_cases "github.com/GuiCezaF/microsservico-go/internal/use-cases"
	"github.com/gin-gonic/gin"
)

func ListCategories(ctx *gin.Context, repository repositories.ICategoryRepository) {
	useCases := use_cases.NewListCategoriesUseCase(repository)

	categories, err := useCases.Execute()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"success": false,
				"error":   err.Error(),
			})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"success": true, "categories": categories})
}
