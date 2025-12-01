package services

import (
	"context"
	"errors"
	"marketplace/internal/store/pgstore"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PaymentMethodsService struct {
	pool    *pgxpool.Pool
	queries *pgstore.Queries
}

func NewPaymentMethodsService(pool *pgxpool.Pool) PaymentMethodsService {
	return PaymentMethodsService{
		pool:    pool,
		queries: pgstore.New(pool),
	}
}

var ErrFailedToGetPaymentMethods = errors.New("failed to get payment methods")

func (pms *PaymentMethodsService) GetPaymentMethods(ctx context.Context) ([]pgstore.PaymentMethod, error) {
	paymentMethods, err := pms.queries.ListPaymentMethods(ctx)
	if err != nil {
		return []pgstore.PaymentMethod{}, ErrFailedToGetPaymentMethods
	}

	return paymentMethods, nil
}
