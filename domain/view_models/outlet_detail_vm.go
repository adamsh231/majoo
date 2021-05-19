package view_models

import "github.com/adamsh231/majoo/domain/models"

type OutletDetailVM struct {
	ID       string   `json:"id"`
	Merchant merchant `json:"merchant"`
	Name     string   `json:"name"`
	Phone    string   `json:"phone"`
	Address  string   `json:"address"`
}

func NewOutletDetailVM() *OutletDetailVM {
	return &OutletDetailVM{}
}

func (vm *OutletDetailVM) Build(model *models.Outlet) {
	*vm = OutletDetailVM{
		ID: model.ID,
		Merchant: merchant{
			ID:   model.Merchant.ID,
			Name: model.Merchant.Name,
		},
		Name:    model.Name,
		Phone:   model.Phone,
		Address: model.Address,
	}
}
