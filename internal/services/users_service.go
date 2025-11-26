package services

import (
	"context"
	"errors"
	"marketplace/internal/store/pgstore"
	"marketplace/internal/usecases/user"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	pool    *pgxpool.Pool
	queries *pgstore.Queries
}

func NewUserService(pool *pgxpool.Pool) UserService {
	return UserService{
		pool:    pool,
		queries: pgstore.New(pool),
	}
}

var (
	ErrUserEmailAlreadyExists = errors.New("email already exists")
	ErrUserTelAlreadyExists   = errors.New("tel already exists")
	ErrFailedToCreateUser     = errors.New("failed to create user")
)

func (us *UserService) CreateUser(ctx context.Context, user user.CreateUserRequest) (uuid.UUID, error) {
	_, err := us.queries.GetUserByEmail(ctx, user.Email)
	if err == nil {
		return uuid.UUID{}, ErrUserEmailAlreadyExists
	}

	_, err = us.queries.GetUserByTel(ctx, user.Tel)
	if err == nil {
		return uuid.UUID{}, ErrUserTelAlreadyExists
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err != nil {
		return uuid.UUID{}, ErrFailedToCreateUser
	}

	createdUser, err := us.queries.CreateUser(ctx, pgstore.CreateUserParams{
		Name:     user.Name,
		Email:    user.Email,
		Tel:      user.Tel,
		Password: string(hashedPassword),
		Avatar: pgtype.Text{
			String: user.Avatar,
			Valid:  user.Avatar != "",
		},
	})
	if err != nil {
		return uuid.UUID{}, ErrFailedToCreateUser
	}

	return createdUser.ID, nil
}
