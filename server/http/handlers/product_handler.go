package handlers

import (
	"github.com/adamsh231/majoo/domain/requests"
	"github.com/adamsh231/majoo/usecase"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type ProductHandler struct {
	Handler
}

func (handler AuthHandler) Add(ctx *fiber.Ctx) (err error) {
	// Parse & Checking input
	input := new(requests.UserAddRequest)
	if err := ctx.BodyParser(input); err != nil {
		return handler.SendResponse(ctx, ResponseWithOutMeta, nil, nil, err, http.StatusBadRequest)
	}
	if err := handler.UcContract.Validate.Struct(input); err != nil {
		return handler.SendResponse(ctx, ResponseWithOutMeta, nil, nil, err, http.StatusBadRequest)
	}

	// Database processing
	handler.UcContract.PostgresTX, err = handler.UcContract.PostgresDB.Begin()
	if err != nil {
		handler.UcContract.PostgresTX.Rollback()
		return handler.SendResponse(ctx, ResponseWithOutMeta, nil, nil, err, http.StatusUnprocessableEntity)
	}
	uc := usecase.NewUserUseCase(handler.UcContract)
	res, err := uc.Add(input)
	if err != nil {
		handler.UcContract.PostgresTX.Rollback()
		return handler.SendResponse(ctx, ResponseWithOutMeta, nil, nil, err, http.StatusUnprocessableEntity)
	}
	handler.UcContract.PostgresTX.Commit()

	return handler.SendResponse(ctx, ResponseWithOutMeta, res, nil, err, http.StatusUnprocessableEntity)
}
