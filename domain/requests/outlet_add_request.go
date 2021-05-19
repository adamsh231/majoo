package requests

type OutletAddRequest struct {
	MerchantID string `json:"merchant_id" validate:"required"`
	Name       string `json:"name" validate:"required"`
	Phone      string `json:"phone" validate:"required"`
	Address    string `json:"address" validate:"required"`
}
