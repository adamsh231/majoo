package interfaces

import (
	"github.com/adamsh231/majoo/domain/requests"
	"github.com/adamsh231/majoo/domain/view_models"
)

type IMerchantUseCase interface {
	Browse(search, orderBy, sort string, page, limit int) (res []*view_models.MerchantListVM, pagination view_models.PaginationVm, err error)

	Read(id string) (res *view_models.MerchantDetailVM, err error)

	Add(req *requests.MerchantAddRequest) (res string, err error)

	Edit(req *requests.MerchantEditRequest, id string) (res string, err error)

	Delete(id string) (res string, err error)
}
