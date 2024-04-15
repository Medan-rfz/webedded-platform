package auth_handler

import (
	"context"

	auth_dto "authorization/internal/domain/dto/auth"
)

func (h *authHandler) Login(ctx context.Context, data auth_dto.LoginDTO) (string, error) {
	return h.authService.Login(ctx, data)
}
