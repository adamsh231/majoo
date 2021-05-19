package usecase

import (
	"github.com/adamsh231/majoo/domain/interfaces"
	"github.com/adamsh231/majoo/domain/requests"
	"github.com/adamsh231/majoo/domain/view_models"
)

type ProductOutletUseCase struct {
	*Contract
}

func NewProductOutletUseCase(ucContract *Contract) interfaces.IProductOutletUseCase {
	return &ProductOutletUseCase{Contract: ucContract}
}

func (uc ProductOutletUseCase) Browse(search, orderBy, sort string, page, limit int) (res []*view_models.ProductOutletListVM, pagination view_models.PaginationVm, err error) {
	panic("implement me")
}

func (uc ProductOutletUseCase) Read(id string) (res view_models.ProductOutletDetailVM, err error) {
	panic("implement me")
}

func (uc ProductOutletUseCase) Add(req *requests.ProductOutletAddRequest) (res string, err error) {
	panic("implement me")
}

func (uc ProductOutletUseCase) Edit(req *requests.ProductOutletEditRequest, id string) (res string, err error) {
	panic("implement me")
}

func (uc ProductOutletUseCase) Delete(id string) (res string, err error) {
	panic("implement me")
}
