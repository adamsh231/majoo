package view_models

import (
	"github.com/adamsh231/majoo/domain/models"
	"strings"
)

type ProductDetailVM struct {
	ID          string   `json:"id"`
	Merchant    merchant `json:"merchant"`
	Sku         string   `json:"sku"`
	Name        string   `json:"name"`
	Slug        string   `json:"slug"`
	Description string   `json:"description"`
	Images      []string `json:"images"`
}

type merchant struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func NewProductDetailVM() *ProductDetailVM {
	return &ProductDetailVM{}
}

func (vm *ProductDetailVM) Build(model models.Product) {
	*vm = ProductDetailVM{
		ID: model.ID,
		Merchant: merchant{
			ID:   model.Merchant.ID,
			Name: model.Merchant.Name,
		},
		Sku:         model.Sku,
		Name:        model.Name,
		Slug:        model.Slug,
		Description: model.Description,
		Images: strings.Split(model.Images, ","),
	}
}
