package view_models

import "github.com/adamsh231/majoo/domain/models"

type ProductImageDetailVM struct {
	ID   string `json:"id"`
	Path string `json:"path"`
	Alt  string `json:"alt"`
}

func NewProductImageDetailVM() *ProductImageDetailVM {
	return &ProductImageDetailVM{}
}

func (vm *ProductImageDetailVM) Build(model models.ProductImage) {
	*vm = ProductImageDetailVM{
		ID:   model.ID,
		Path: model.Path,
		Alt:  model.Alt,
	}
}
