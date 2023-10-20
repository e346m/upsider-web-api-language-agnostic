package http

import (
	usecases "github.com/e346m/upsider-wala/internal/usecases"
)

type Handler struct {
	usecase usecases.Usecase
}

func NewHandler(usecase *usecases.Usecase) *Handler {
	return &Handler{
		usecase: *usecase,
	}
}
