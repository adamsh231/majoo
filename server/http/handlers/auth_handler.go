package handlers

import (
	"github.com/adamsh231/majoo/usecase"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type AuthHandler struct {
	Handler
}

func (handler AuthHandler) Login(ctx *fiber.Ctx) error {
	email := "adamsyarif219@gmail.com"
	pass := "ujseag"


	uc := usecase.NewAuthUseCase(handler.UcContract)
	res, err := uc.Login(email, pass)

	return handler.SendResponse(ctx, ResponseWithOutMeta, res, nil, err, http.StatusUnprocessableEntity)
}

