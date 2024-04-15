package auth_http

import (
	"context"

	auth_dto "authorization/internal/domain/dto/auth"
	authdto "authorization/internal/domain/dto/auth"
)

type authHandler interface {
	Login(ctx context.Context, data auth_dto.LoginDTO) (string, error)
	Register(ctx context.Context, data authdto.RegisterDTO) error
}

type authHttpHandler struct {
	authHandler authHandler
}

func NewAuthHttpHandler(authHandler authHandler) *authHttpHandler {
	return &authHttpHandler{
		authHandler: authHandler,
	}
}
