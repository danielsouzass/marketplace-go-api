package product

import "marketplace/internal/store/pgstore"

type ProductResponse struct {
	pgstore.Product
	Images         []pgstore.ProductImage  `json:"images"`
	PaymentMethods []pgstore.PaymentMethod `json:"payment_methods"`
}

type GetProductsResponse = []ProductResponse
