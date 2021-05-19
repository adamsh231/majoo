package view_models

import "github.com/adamsh231/majoo/domain/models"

type ProductImageListVM struct {
	ID   string `json:"id"`
	Path string `json:"path"`
}

func NewProductImageListVM() *ProductImageListVM {
	return &ProductImageListVM{}
}

func (vm *ProductImageListVM) Build(model models.ProductImage) {
	*vm = ProductImageListVM{
		ID:   model.ID,
		Path: model.Path,
	}
}
