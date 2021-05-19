package view_models

import "github.com/adamsh231/majoo/domain/models"

type OutletListVM struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func NewOutletListVM() *OutletListVM {
	return &OutletListVM{}
}

func (vm *OutletListVM) Build(model *models.Outlet) {
	*vm = OutletListVM{
		ID:    model.ID,
		Name:  model.Name,
	}
}
