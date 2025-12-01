package services

import (
	"context"
	"errors"
	"marketplace/internal/store/pgstore"
	"marketplace/internal/usecases/product"

	"github.com/go-chi/jwtauth/v5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductService struct {
	pool    *pgxpool.Pool
	queries *pgstore.Queries
}

func NewProductService(pool *pgxpool.Pool) ProductService {
	return ProductService{
		pool:    pool,
		queries: pgstore.New(pool),
	}
}

var (
	ErrFailedToCreateProduct            = errors.New("failed to create product")
	ErrPriceInvalid                     = errors.New("price is invalid")
	ErrFailedToCreateProductImage       = errors.New("failed to create product image")
	ErrFailedToFindPaymentMethods       = errors.New("failed to find payment methods")
	ErrInvalidPaymentMethods            = errors.New("one or more payment methods are invalid")
	ErrFailedToCreateProductPayment     = errors.New("failed to create product payment method")
	ErrFailedToGetProducts              = errors.New("failed to get products")
	ErrFailedToGetProductImages         = errors.New("failed to get product images")
	ErrFailedToGetProductPaymentMethods = errors.New("failed to get product payment methods")
)

func (ps *ProductService) CreateProduct(ctx context.Context, data product.CreateProductRequest) (uuid.UUID, error) {
	_, claims, _ := jwtauth.FromContext(ctx)

	authenticatedUserId, err := uuid.Parse(claims["user_id"].(string))
	if err != nil {
		return uuid.UUID{}, ErrUserNotFound
	}

	userFound, err := ps.queries.GetUserByID(ctx, authenticatedUserId)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return uuid.UUID{}, ErrUserNotFound
		}
		return uuid.UUID{}, ErrUserNotFound
	}

	validPaymentMethods, err := ps.queries.GetPaymentMethodsByKeys(ctx, data.PaymentMethods)
	if err != nil {
		return uuid.UUID{}, ErrFailedToFindPaymentMethods
	}

	if len(validPaymentMethods) != len(data.PaymentMethods) {
		return uuid.UUID{}, ErrInvalidPaymentMethods
	}

	createdProduct, err := ps.queries.CreateProduct(ctx, pgstore.CreateProductParams{
		UserID:      userFound.ID,
		Name:        data.Name,
		Description: data.Description,
		IsNew:       data.IsNew,
		Price:       data.Price,
		AcceptTrade: data.AcceptTrade,
	})
	if err != nil {
		return uuid.UUID{}, ErrFailedToCreateProduct
	}

	for _, productImageURL := range data.Images {
		_, err := ps.queries.CreateProductImage(ctx, pgstore.CreateProductImageParams{
			ProductID: createdProduct.ID,
			Path:      productImageURL,
		})
		if err != nil {
			return uuid.UUID{}, ErrFailedToCreateProductImage
		}
	}

	findPaymentMethodIDByKey := func(key string) (uuid.UUID, error) {
		for _, pm := range validPaymentMethods {
			if pm.Key == key {
				return pm.ID, nil
			}
		}
		return uuid.UUID{}, ErrInvalidPaymentMethods
	}

	for _, paymentMethodKey := range data.PaymentMethods {
		paymentMethodID, err := findPaymentMethodIDByKey(paymentMethodKey)
		if err != nil {
			return uuid.UUID{}, err
		}

		_, err = ps.queries.CreateProductPaymentMethod(ctx, pgstore.CreateProductPaymentMethodParams{
			ProductID:       createdProduct.ID,
			PaymentMethodID: paymentMethodID,
		})
		if err != nil {
			return uuid.UUID{}, ErrFailedToCreateProductPayment
		}
	}

	return createdProduct.ID, nil
}

func (ps *ProductService) GetProducts(ctx context.Context) (product.GetProductsResponse, error) {
	_, claims, _ := jwtauth.FromContext(ctx)
	userId, err := uuid.Parse(claims["user_id"].(string))
	if err != nil {
		return product.GetProductsResponse{}, ErrUserNotFound
	}

	_, err = ps.queries.GetUserByID(ctx, userId)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return product.GetProductsResponse{}, ErrUserNotFound
		}
		return product.GetProductsResponse{}, ErrUserNotFound
	}

	dbProducts, err := ps.queries.GetProductsByUserID(ctx, userId)
	if err != nil {
		return product.GetProductsResponse{}, ErrFailedToGetProducts
	}

	products := make([]product.ProductResponse, 0, len(dbProducts))

	for _, p := range dbProducts {
		images, err := ps.queries.GetProductImagesByProductID(ctx, p.ID)
		if err != nil {
			return product.GetProductsResponse{}, ErrFailedToGetProductImages
		}

		paymentMethods, err := ps.queries.GetProductPaymentMethodsByProductID(ctx, p.ID)
		if err != nil {
			return product.GetProductsResponse{}, ErrFailedToGetProductPaymentMethods
		}

		productResponse := product.ProductResponse{
			Product:        p,
			Images:         images,
			PaymentMethods: paymentMethods,
		}

		products = append(products, productResponse)
	}

	return products, nil
}
