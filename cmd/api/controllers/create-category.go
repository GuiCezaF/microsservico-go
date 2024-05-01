package controllers

import (
	"net/http"

	"github.com/GuiCezaF/microsservico-go/internal/repositories"
	use_cases "github.com/GuiCezaF/microsservico-go/internal/use-cases"
	"github.com/gin-gonic/gin"
)

type createCategoryInput struct {
	Name string `json:"name" binding:"required"`
}

func CreateCategory(ctx *gin.Context, repository repositories.ICategoryRepository) {
	var body createCategoryInput
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"success": false,
				"error":   err.Error(),
			})
		return
	}

	useCases := use_cases.NewCreateCategoryUseCase(repository)

	err := useCases.Execute(body.Name)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"success": false,
				"error":   err.Error(),
			})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"success": true})
}
