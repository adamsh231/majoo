package interfaces

import (
	"github.com/adamsh231/majoo/domain/requests"
	"github.com/adamsh231/majoo/domain/view_models"
)

type IProductOutletUseCase interface {
	Browse(search, orderBy, sort string, page, limit int) (res []*view_models.ProductOutletListVM, pagination view_models.PaginationVm, err error)

	Read(id string) (res view_models.ProductOutletDetailVM, err error)

	Add(req *requests.ProductOutletAddRequest) (res string, err error)

	Edit(req *requests.ProductOutletEditRequest, id string) (res string, err error)

	Delete(id string) (res string, err error)
}
