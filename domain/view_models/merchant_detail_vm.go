package view_models

import "github.com/adamsh231/majoo/domain/models"

type MerchantDetailVM struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

func NewMerchantDetailVM() *MerchantDetailVM {
	return &MerchantDetailVM{}
}

func (vm *MerchantDetailVM) Build(model *models.Merchant) {
	*vm = MerchantDetailVM{
		ID:      model.ID,
		Name:    model.Name,
		Phone:   model.Phone,
		Address: model.Address,
	}
}
