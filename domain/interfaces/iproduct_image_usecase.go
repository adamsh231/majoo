package interfaces

import (
	"github.com/adamsh231/majoo/domain/view_models"
)

type IProductImageUseCase interface {
	Browse(search, orderBy, sort string, page, limit int) (res []*view_models.ProductImageListVM, pagination view_models.PaginationVm, err error)

	Read(id string) (res view_models.ProductImageDetailVM, err error)

	Add() (res string, err error)

	Edit() (res string, err error)

	Delete(id string) (res string, err error)
}