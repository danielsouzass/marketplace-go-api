package product

type CreateProductRequest struct {
	Name           string   `json:"name" validate:"required"`
	Description    string   `json:"description" validate:"required"`
	IsNew          bool     `json:"is_new" validate:"required"`
	Price          float64  `json:"price" validate:"required"`
	AcceptTrade    bool     `json:"accept_trade" validate:"required"`
	Images         []string `json:"images" validate:"required,min=1,max=3,dive,url"`
	PaymentMethods []string `json:"payment_methods" validate:"required,min=1,dive,required,oneof=boleto pix cash deposit card"`
}

type CreateProductResponse struct {
	Id string `json:"id"`
}
