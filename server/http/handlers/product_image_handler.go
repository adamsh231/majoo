package handlers

import (
	"errors"
	"github.com/adamsh231/majoo/domain/requests"
	"github.com/adamsh231/majoo/packages/messages"
	"github.com/adamsh231/majoo/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"net/http"
	"path/filepath"
)

type ProductImageHandler struct {
	Handler
}

func (handler ProductImageHandler) Add(ctx *fiber.Ctx) (err error) {
	//get body form file and form value
	file, err := ctx.FormFile("image")
	if err != nil {
		return handler.SendResponse(ctx, ResponseWithOutMeta, nil, nil, err, http.StatusUnprocessableEntity)
	}

	fileType := filepath.Ext(file.Filename)
	if !handler.CheckImageType(fileType) {
		return handler.SendResponse(ctx, ResponseWithOutMeta, nil, nil, errors.New(messages.InvalidImageType), http.StatusUnprocessableEntity)
	}

	fileName := uuid.New().String() + fileType
	err = ctx.SaveFile(file, handler.UcContract.ImageDirectory+fileName)
	if err != nil {
		return handler.SendResponse(ctx, ResponseWithOutMeta, nil, nil, err, http.StatusUnprocessableEntity)
	}

	input := requests.ProductImageAddRequest{
		ProductID: ctx.FormValue("product_id"),
		Alt:       ctx.FormValue("alt"),
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
	uc := usecase.NewProductImageUseCase(handler.UcContract)
	res, err := uc.Add(&input, fileName)
	if err != nil {
		handler.UcContract.PostgresTX.Rollback()
		return handler.SendResponse(ctx, ResponseWithOutMeta, nil, nil, err, http.StatusUnprocessableEntity)
	}
	handler.UcContract.PostgresTX.Commit()

	return handler.SendResponse(ctx, ResponseWithOutMeta, res, nil, err, http.StatusUnprocessableEntity)

	return nil
}
