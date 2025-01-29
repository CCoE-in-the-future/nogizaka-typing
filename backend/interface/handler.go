package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"backend/domain/entities"
	usecase "backend/usecase/generation"
)

// type GenerationHandler interface {
// 	GetAllGenerations(c echo.Context) error
// }

type GenerationHandler struct {
	generationUseCase usecase.GenerationUseCase
}

func NewGenerationHandler(guc usecase.GenerationUseCase) *GenerationHandler {
	return &GenerationHandler{generationUseCase: guc}
}

func (h *GenerationHandler) GetAllGenerations(c echo.Context) error {
	// generations, err := h.generationUseCase.GetAllGenerations()
	// if err != nil {
	// 	return c.JSON(http.StatusInternalServerError, map[string]string{
	// 		"error": "failed to fetch generations",
	// 	})
	// }

	generations := []entities.Generation{
		{Generation: "6th", Image: "Unknown"},
	}
	return c.JSON(http.StatusOK, generations)
}
