package view_models

import "github.com/adamsh231/majoo/domain/models"

type ProductOutletDetailVM struct {
	ID      string  `json:"id"`
	Outlet  outlet  `json:"outlet"`
	Product product `json:"product"`
	Price   float64 `json:"price"`
	Stock   int     `json:"stock"`
}

type outlet struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type product struct {
	ID   string `json:id`
	Sku  string `json:"sku"`
	Name string `json:"name"`
}

func NewProductOutletDetailVM() *ProductOutletDetailVM {
	return &ProductOutletDetailVM{}
}

func (vm *ProductOutletDetailVM) Build(model models.ProductOutlet) {
	*vm = ProductOutletDetailVM{
		ID: model.ID,
		Outlet: outlet{
			ID:   model.Outlet.ID,
			Name: model.Outlet.Name,
		},
		Product: product{
			ID:   model.Product.ID,
			Sku:  model.Product.Sku,
			Name: model.Product.Name,
		},
		Price: model.Price,
		Stock: model.Stock,
	}
}
