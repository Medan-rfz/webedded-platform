package auth_http

import (
	"context"

	auth_dto "webedded.users_management/internal/domain/dto/auth"
)

type authHandler interface {
	Login(ctx context.Context, data auth_dto.LoginDTO) (*auth_dto.LoginSuccess, error)
	Register(ctx context.Context, data auth_dto.RegisterDTO) error
	Refresh(ctx context.Context, data auth_dto.RefreshDTO) (string, error)
}

type authHttpHandler struct {
	authHandler authHandler
}

func NewAuthHttpHandler(authHandler authHandler) *authHttpHandler {
	return &authHttpHandler{
		authHandler: authHandler,
	}
}
