package handlers

import (
	"github.com/adamsh231/majoo/domain/requests"
	"github.com/adamsh231/majoo/usecase"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
)

type MerchantHandler struct {
	Handler
}

func (handler MerchantHandler) Browse(ctx *fiber.Ctx) (err error) {
	// Get Query Param
	search := ctx.Query("search")
	orderBy := ctx.Query("order_by")
	sort := ctx.Query("sort")
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	page, _ := strconv.Atoi(ctx.Query("page"))

	// Database Processing
	uc := usecase.NewMerchantUseCase(handler.UcContract)
	res, pagination, err := uc.Browse(search, orderBy, sort, page, limit)

	return handler.SendResponse(ctx, ResponseWithMeta, res, pagination, err, http.StatusUnprocessableEntity)
}

func (handler MerchantHandler) Add(ctx *fiber.Ctx) (err error) {
	// Parse & Checking input
	input := new(requests.MerchantAddRequest)
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
	uc := usecase.NewMerchantUseCase(handler.UcContract)
	res, err := uc.Add(input)
	if err != nil {
		handler.UcContract.PostgresTX.Rollback()
		return handler.SendResponse(ctx, ResponseWithOutMeta, nil, nil, err, http.StatusUnprocessableEntity)
	}
	handler.UcContract.PostgresTX.Commit()

	return handler.SendResponse(ctx, ResponseWithOutMeta, res, nil, err, http.StatusUnprocessableEntity)
}

