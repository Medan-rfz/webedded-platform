package auth_handler

import (
	"context"

	auth_dto "authorization/internal/domain/dto/auth"
	"authorization/internal/domain/entities/users"
)

type authService interface {
	Login(ctx context.Context, data auth_dto.LoginDTO) (string, error)
	Register(ctx context.Context, data auth_dto.AuthRegisterDTO) error
	CheckUserExists(ctx context.Context, data auth_dto.RegisterDTO) (bool, error)
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
