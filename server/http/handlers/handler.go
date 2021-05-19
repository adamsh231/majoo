package handlers

import (
	"github.com/adamsh231/majoo/domain/view_models"
	"github.com/adamsh231/majoo/packages/helper"
	"github.com/adamsh231/majoo/usecase"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strings"
)

type Handler struct {
	UcContract *usecase.Contract
}

const (
	ResponseWithMeta    = `responseWithMeta`
	ResponseWithOutMeta = `responseWithOutMeta`
)

func (handler Handler) SendResponse(ctx *fiber.Ctx, responseType string, data interface{}, meta interface{}, err interface{}, statusCode int) error {
	if err != nil {
		if statusCode == 400 {
			errorMessage := err.(validator.ValidationErrors)
			return handler.ResponseValidationError(ctx, errorMessage)
		} else {
			return handler.SendErrorResponse(ctx, err.(error).Error(), statusCode)
		}
	} else {
		if responseType == ResponseWithMeta {
			return handler.SendSuccessResponseWithMeta(ctx, data, meta)
		} else {
			return handler.SendSuccessResponseWithOutMeta(ctx, data)
		}
	}
}

func (handler Handler) SendSuccessResponseWithOutMeta(ctx *fiber.Ctx, data interface{}) error {
	return ctx.Status(http.StatusOK).JSON(view_models.SuccessResponseWithOutMetaVm{Data: data})
}

func (handler Handler) SendSuccessResponseWithMeta(ctx *fiber.Ctx, data interface{}, meta interface{}) error {
	return ctx.Status(http.StatusOK).JSON(view_models.SuccessResponseWithMeta{Data: data, Meta: meta})
}

func (handler Handler) SendErrorResponse(ctx *fiber.Ctx, err interface{}, statusCode int) error {
	return ctx.Status(statusCode).JSON(view_models.ErrorResponseVm{Message: err})
}

func (handler Handler) ResponseValidationError(ctx *fiber.Ctx, error validator.ValidationErrors) error {
	errorMessage := handler.ExtractErrorValidationMessages(error)

	return handler.SendErrorResponse(ctx, errorMessage, http.StatusBadRequest)
}

func (handler Handler) ExtractErrorValidationMessages(error validator.ValidationErrors) map[string][]string {
	errorMessage := map[string][]string{}
	errorTranslation := error.Translate(handler.UcContract.Translator)

	for _, err := range error {
		errKey := helper.Underscore(err.StructField())
		errorMessage[errKey] = append(
			errorMessage[errKey],
			strings.Replace(errorTranslation[err.Namespace()], err.StructField(), err.StructField(), -1),
		)
	}

	return errorMessage
}

func (handler Handler) CheckImageType(imageType string) bool {
	isValid := false
	imageTypes := []string{".png", ".jpg", "jpeg"}
	for _, valImageType := range imageTypes {
		if valImageType == strings.ToLower(imageType) {
			isValid = true
			break
		}
	}

	return isValid
}
