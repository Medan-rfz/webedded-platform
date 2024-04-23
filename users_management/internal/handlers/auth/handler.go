package auth_handler

import (
	"context"

	"users_management/internal/domain/entities/users"

	auth_dto "users_management/internal/domain/dto/auth"
)

type authService interface {
	Login(ctx context.Context, data auth_dto.LoginDTO) (*auth_dto.LoginSuccess, error)
	Register(ctx context.Context, data auth_dto.RegisterDTO) error
	Refresh(ctx context.Context, data auth_dto.RefreshDTO) (string, error)
}

type usersService interface {
	Create(ctx context.Context, user auth_dto.RegisterDTO) (*users.User, error)
	DeleteById(ctx context.Context, userId int64) error
}

type authHandler struct {
	authService  authService
	usersService usersService
}

func NewAuthHandler(authService authService, usersService usersService) *authHandler {
	return &authHandler{
		authService:  authService,
		usersService: usersService,
	}
}
