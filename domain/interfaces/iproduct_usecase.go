package interfaces

import (
	"github.com/adamsh231/majoo/domain/requests"
	"github.com/adamsh231/majoo/domain/view_models"
)

type IProductUseCase interface {
	Browse(search, orderBy, sort string, page, limit int) (res []*view_models.ProductListVM, pagination view_models.PaginationVm, err error)

	Read(id string) (res *view_models.ProductDetailVM, err error)

	Add(req *requests.ProductAddRequest) (res string, err error)

	Edit(req *requests.ProductEditRequest, id string) (res string, err error)

	Delete(id string) (res string, err error)
}
