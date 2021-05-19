package usecase

import (
	"github.com/adamsh231/majoo/domain/interfaces"
	"github.com/adamsh231/majoo/domain/requests"
	"github.com/adamsh231/majoo/domain/view_models"
)

type MerchantUseCase struct {
	*Contract
}

func NewMerchantUseCase(ucContract *Contract) interfaces.IMerchantUseCase{
	return &MerchantUseCase{Contract: ucContract}
}

func (m MerchantUseCase) Browse(search, orderBy, sort string, page, limit int) (res []*view_models.MerchantListVM, pagination view_models.PaginationVm, err error) {
	panic("implement me")
}

func (m MerchantUseCase) Read(id string) (res view_models.MerchantDetailVM, err error) {
	panic("implement me")
}

func (m MerchantUseCase) Add(req *requests.MerchantAddRequest) (res string, err error) {
	panic("implement me")
}

func (m MerchantUseCase) Edit(req *requests.MerchantEditRequest, id string) (res string, err error) {
	panic("implement me")
}

func (m MerchantUseCase) Delete(id string) (res string, err error) {
	panic("implement me")
}