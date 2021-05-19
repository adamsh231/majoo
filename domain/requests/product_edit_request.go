package requests

type ProductEditRequest struct {
	MerchantID  string `json:"merchant_id" validate:"required"`
	Sku         string `json:"sku" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Slug        string `json:"slug" validate:"required"`
	Description string `json:"description" validate:"required"`
}
