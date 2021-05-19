package view_models

import "github.com/adamsh231/majoo/domain/models"

type ProductListVM struct {
	ID   string `json:"id"`
	Sku  string `json:"sku"`
	Name string `json:"name"`
}

func NewProductListVM() *ProductListVM {
	return &ProductListVM{}
}

func (vm *ProductListVM) Build(model *models.Product) {
	*vm = ProductListVM{
		ID:   model.ID,
		Sku:  model.Sku,
		Name: model.Name,
	}
}
