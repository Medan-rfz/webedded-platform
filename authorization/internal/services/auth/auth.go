package auth_service

import (
	"context"

	auth_entities "authorization/internal/domain/entities/auth"
)

var jwtKey = []byte("secret-key")

type usersRepo interface {
	GetById(ctx context.Context, id int64) (*auth_entities.User, error)
	GetByEmail(ctx context.Context, email string) (*auth_entities.User, error)
	Create(ctx context.Context, user auth_entities.User) error
}

type authService struct {
	usersRepo usersRepo
}

func NewUsersService(usersRepo usersRepo) *authService {
	return &authService{
		usersRepo: usersRepo,
	}
}
