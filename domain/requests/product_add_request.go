package requests

type ProductAddRequest struct {
	MerchantID  string `json:"merchant_id" validate:"required"`
	Sku         string `json:"sku" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
}
