package view_models

import "github.com/adamsh231/majoo/domain/models"

type MerchantListVM struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func NewMerchantListVM() *MerchantListVM {
	return &MerchantListVM{}
}

func (vm *MerchantListVM) Build(model *models.Merchant) {
	*vm = MerchantListVM{
		ID:    model.ID,
		Name:  model.Name,
	}
}
