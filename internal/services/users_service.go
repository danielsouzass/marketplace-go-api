package services

import (
	"context"
	"errors"
	"marketplace/internal/auth"
	"marketplace/internal/store/pgstore"
	"marketplace/internal/usecases/user"
	"time"

	"github.com/go-chi/jwtauth/v5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
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
	ErrInvalidCredentials     = errors.New("email or password is incorrect")
	ErrFailedToAuthenticate   = errors.New("failed to authenticate user")
	ErrUserNotFound           = errors.New("user not found")
	ErrFailedToRefreshToken   = errors.New("failed to refresh token")
	ErrRefreshTokenNotFound   = errors.New("refresh token not found")
	ErrRefreshTokenExpired    = errors.New("refresh token has expired")
	ErrInvalidRefreshToken    = errors.New("invalid refresh token")
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

func (us *UserService) AuthenticateUser(ctx context.Context, credentials user.LoginUserRequest) (user.LoginResponse, error) {
	userFound, err := us.queries.GetUserByEmail(ctx, credentials.Email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return user.LoginResponse{}, ErrInvalidCredentials
		}
		return user.LoginResponse{}, ErrFailedToAuthenticate
	}

	err = bcrypt.CompareHashAndPassword([]byte(userFound.Password), []byte(credentials.Password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return user.LoginResponse{}, ErrInvalidCredentials
		}
		return user.LoginResponse{}, ErrFailedToAuthenticate
	}

	accessToken, err := us.GenerateAccessToken(userFound.ID)
	if err != nil {
		return user.LoginResponse{}, err
	}

	refreshToken, err := us.GenerateRefreshToken(ctx, userFound.ID)
	if err != nil {
		return user.LoginResponse{}, err
	}

	return user.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (us *UserService) GetUser(ctx context.Context) (user.GetUserResponse, error) {
	_, claims, _ := jwtauth.FromContext(ctx)

	userId, err := uuid.Parse(claims["user_id"].(string))
	if err != nil {
		return user.GetUserResponse{}, ErrUserNotFound
	}

	userFound, err := us.queries.GetUserByID(ctx, userId)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return user.GetUserResponse{}, ErrUserNotFound
		}
		return user.GetUserResponse{}, ErrUserNotFound
	}

	return user.GetUserResponse{
		ID:        userFound.ID.String(),
		Name:      userFound.Name,
		Email:     userFound.Email,
		Tel:       userFound.Tel,
		Avatar:    userFound.Avatar.String,
		CreatedAt: userFound.CreatedAt,
		UpdatedAt: userFound.UpdatedAt,
	}, nil
}

func (us *UserService) RefreshUserToken(ctx context.Context, refreshToken string) (user.RefreshTokenUserResponse, error) {
	refreshTokenId, err := uuid.Parse(refreshToken)
	if err != nil {
		return user.RefreshTokenUserResponse{}, ErrInvalidRefreshToken
	}

	refreshTokenFound, err := us.queries.GetRefreshTokenByID(ctx, refreshTokenId)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return user.RefreshTokenUserResponse{}, ErrRefreshTokenNotFound
		}
		return user.RefreshTokenUserResponse{}, ErrFailedToRefreshToken
	}

	refreshTokenExpired := time.Now().Unix() > int64(refreshTokenFound.ExpiresIn)

	if refreshTokenExpired {
		return user.RefreshTokenUserResponse{}, ErrRefreshTokenExpired
	}

	newAccessToken, err := us.GenerateAccessToken(refreshTokenFound.UserID)
	if err != nil {
		return user.RefreshTokenUserResponse{}, err
	}

	newRefreshToken, err := us.GenerateRefreshToken(ctx, refreshTokenFound.UserID)
	if err != nil {
		return user.RefreshTokenUserResponse{}, err
	}

	return user.RefreshTokenUserResponse{
		AccessToken:  newAccessToken,
		RefreshToken: newRefreshToken,
	}, nil
}

func (us *UserService) GenerateAccessToken(userID uuid.UUID) (accessToken string, err error) {
	newAccessToken, err := auth.NewAccessToken(userID)
	if err != nil {
		return "", ErrFailedToAuthenticate
	}
	return newAccessToken, nil
}

func (us *UserService) GenerateRefreshToken(ctx context.Context, userID uuid.UUID) (refreshToken string, err error) {
	oldRefreshToken, err := us.queries.GetRefreshTokenByUserID(ctx, userID)
	if err == nil {
		err := us.queries.DeleteRefreshToken(ctx, oldRefreshToken.ID)
		if err != nil {
			return "", ErrFailedToAuthenticate
		}
	}

	newRefreshToken, err := us.queries.CreateRefreshToken(ctx, pgstore.CreateRefreshTokenParams{
		ExpiresIn: int32(time.Now().Add(24 * time.Hour).Unix()),
		UserID:    userID,
	})
	if err != nil {
		return "", ErrFailedToAuthenticate
	}

	return newRefreshToken.ID.String(), nil
}
