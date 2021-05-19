package usecase

import (
	"github.com/adamsh231/majoo/domain/interfaces"
	"github.com/adamsh231/majoo/domain/view_models"
)

type ProductImageUseCase struct {
	*Contract
}

func NewProductImageUseCase(ucContract *Contract) interfaces.IProductImageUseCase {
	return &ProductImageUseCase{Contract: ucContract}
}

func (uc ProductImageUseCase) Browse(search, orderBy, sort string, page, limit int) (res []*view_models.ProductImageListVM, pagination view_models.PaginationVm, err error) {
	panic("implement me")
}

func (uc ProductImageUseCase) Read(id string) (res view_models.ProductImageDetailVM, err error) {
	panic("implement me")
}

func (uc ProductImageUseCase) Add() (res string, err error) {
	panic("implement me")
}

func (uc ProductImageUseCase) Edit() (res string, err error) {
	panic("implement me")
}

func (uc ProductImageUseCase) Delete(id string) (res string, err error) {
	panic("implement me")
}