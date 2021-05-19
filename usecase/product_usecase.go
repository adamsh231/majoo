package usecase

import (
	"github.com/adamsh231/majoo/domain/interfaces"
	"github.com/adamsh231/majoo/domain/models"
	"github.com/adamsh231/majoo/domain/requests"
	"github.com/adamsh231/majoo/domain/view_models"
	"github.com/adamsh231/majoo/packages/helper"
	"github.com/adamsh231/majoo/repositories"
	"github.com/gosimple/slug"
	"time"
)

type ProductUseCase struct {
	*Contract
}

func NewProductUseCase(ucContract *Contract) interfaces.IProductUseCase {
	return &ProductUseCase{Contract: ucContract}
}

func (uc ProductUseCase) Browse(search, orderBy, sort string, page, limit int) (res []*view_models.ProductListVM, pagination view_models.PaginationVm, err error) {
	repository := repositories.NewProductRepository(uc.PostgresDB)
	offset, limit, page, orderBy, sort := uc.setPaginationParameter(page, limit, orderBy, sort)

	products, err := repository.Browse(search, orderBy, sort, limit, offset)
	if err != nil {
		helper.LogOnly(err.Error(), "uc-browse")
		return res, pagination, err
	}
	for _, product := range products {
		vm := view_models.NewProductListVM()
		vm.Build(&product)
		res = append(res, vm)
	}

	//set pagination
	totalCount, err := repository.Count(search)
	if err != nil {
		helper.LogOnly(err.Error(), "uc-browse-count")
		return res, pagination, err
	}
	pagination = uc.setPaginationResponse(page, limit, totalCount)

	return res, pagination, err
}

func (uc ProductUseCase) Read(id string) (res view_models.ProductDetailVM, err error) {
	panic("implement me")
}

func (uc ProductUseCase) Add(req *requests.ProductAddRequest) (res string, err error) {
	now := time.Now().UTC()
	model := models.Product{
		Merchant:  models.Merchant{ID: req.MerchantID},
		Sku:       req.Sku,
		Name:      req.Name,
		Slug:      slug.Make(req.Name),
		CreatedAt: now,
		UpdatedAt: now,
	}

	repo := repositories.NewProductRepository(uc.PostgresDB)
	res, err = repo.Add(model, uc.PostgresTX)
	if err != nil {
		helper.LogOnly(err.Error(), "uc-add")
		return res, err
	}

	return res, err
}

func (uc ProductUseCase) Edit(req *requests.ProductEditRequest, id string) (res string, err error) {
	panic("implement me")
}

func (uc ProductUseCase) Delete(id string) (res string, err error) {
	panic("implement me")
}
