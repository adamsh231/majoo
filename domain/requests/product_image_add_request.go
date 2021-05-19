package requests

type ProductImageAddRequest struct {
	ProductID string `json:"product_id" validate:"required"`
	Alt       string `json:"alt" validate:"required"`
}
