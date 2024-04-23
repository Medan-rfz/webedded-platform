package auth_handler

import (
	"context"

	auth_dto "users_management/internal/domain/dto/auth"
)

func (h *authHandler) Login(ctx context.Context, data auth_dto.LoginDTO) (*auth_dto.LoginSuccess, error) {
	return h.authService.Login(ctx, data)
}
