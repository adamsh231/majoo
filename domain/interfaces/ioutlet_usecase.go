package interfaces

import (
	"github.com/adamsh231/majoo/domain/requests"
	"github.com/adamsh231/majoo/domain/view_models"
)

type IOutletUseCase interface {
	Browse(search, orderBy, sort string, page, limit int) (res []*view_models.OutletListVM, pagination view_models.PaginationVm, err error)

	Read(id string) (res view_models.OutletDetailVM, err error)

	Add(req *requests.OutletAddRequest) (res string, err error)

	Edit(req *requests.OutletEditRequest, id string) (res string, err error)

	Delete(id string) (res string, err error)
}
