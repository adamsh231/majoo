package usecase

import (
	"github.com/adamsh231/majoo/domain/interfaces"
	"github.com/adamsh231/majoo/domain/requests"
	"github.com/adamsh231/majoo/domain/view_models"
)

type OutletUseCase struct {
	*Contract
}

func NewOutletUseCase(ucContract *Contract) interfaces.IOutletUseCase {
	return &OutletUseCase{Contract: ucContract}
}

func (uc OutletUseCase) Browse(search, orderBy, sort string, page, limit int) (res []*view_models.OutletListVM, pagination view_models.PaginationVm, err error) {
	panic("implement me")
}

func (uc OutletUseCase) Read(id string) (res view_models.OutletDetailVM, err error) {
	panic("implement me")
}

func (uc OutletUseCase) Add(req *requests.OutletAddRequest) (res string, err error) {
	panic("implement me")
}

func (uc OutletUseCase) Edit(req *requests.OutletEditRequest, id string) (res string, err error) {
	panic("implement me")
}

func (uc OutletUseCase) Delete(id string) (res string, err error) {
	panic("implement me")
}
