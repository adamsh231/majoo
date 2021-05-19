package view_models

import "github.com/adamsh231/majoo/domain/models"

type ProductOutletListVM struct {
	ID         string `json:"id"`
	OutletName string `json:"outlet_name"`
	Sku        string `json:"sku"`
	Name       string `json:"name"`
}

func NewProductOutletListVM() *ProductOutletListVM {
	return &ProductOutletListVM{}
}

func (vm *ProductOutletListVM) Build(model models.ProductOutlet) {
	*vm = ProductOutletListVM{
		ID:         model.ID,
		OutletName: model.Outlet.Name,
		Sku:        model.Product.Sku,
		Name:       model.Product.Name,
	}
}
