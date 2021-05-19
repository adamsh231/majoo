package requests

type ProductOutletEditRequest struct {
	ProductID string `json:"product_id" validate:"required"`
	OutletID  string `json:"outlet_id" validate:"required"`
	Price     string `json:"price" validate:"required"`
	Stock     string `json:"stock" validate:"required"`
}
