package handlers

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type TestHandler struct {
	Handler
}

func (handler TestHandler) TestProcess(ctx *fiber.Ctx) (err error){
	return handler.SendResponse(ctx, ResponseWithOutMeta, nil, nil, err, http.StatusUnprocessableEntity)
}
