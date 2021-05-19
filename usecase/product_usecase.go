package usecase

import (
	"github.com/adamsh231/majoo/domain/interfaces"
	"github.com/adamsh231/majoo/domain/requests"
	"github.com/adamsh231/majoo/domain/view_models"
)

type ProductUseCase struct {
	*Contract
}

func NewProductUseCase(ucContract *Contract) interfaces.IProductUseCase {
	return &ProductUseCase{Contract: ucContract}
}

func (uc ProductUseCase) Browse(search, orderBy, sort string, page, limit int) (res []*view_models.ProductListVM, pagination view_models.PaginationVm, err error) {
	panic("implement me")
}

func (uc ProductUseCase) Read(id string) (res view_models.ProductDetailVM, err error) {
	panic("implement me")
}

func (uc ProductUseCase) Add(req *requests.ProductAddRequest) (res string, err error) {
	panic("implement me")
}

func (uc ProductUseCase) Edit(req *requests.ProductEditRequest, id string) (res string, err error) {
	panic("implement me")
}

func (uc ProductUseCase) Delete(id string) (res string, err error) {
	panic("implement me")
}