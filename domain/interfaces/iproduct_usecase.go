package interfaces

import (
	"github.com/adamsh231/majoo/domain/requests"
	"github.com/adamsh231/majoo/domain/view_models"
)

type IProductUseCase interface {
	Browse(search, orderBy, sort string, page, limit int) (res []*view_models.UserListVM, pagination view_models.PaginationVm, err error)

	Read(req *requests.UserAddRequest, id string) (res view_models.LoginVm, err error)

	Add(req *requests.UserAddRequest) (res string, err error)

	Edit(req *requests.UserEditRequest, id string) (res string, err error)

	Delete(id string) (res string, err error)
}