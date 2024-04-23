package auth_handler

import (
	"context"

	auth_dto "users_management/internal/domain/dto/auth"
)

func (h *authHandler) Register(ctx context.Context, data auth_dto.RegisterDTO) error {
	return h.authService.Register(ctx, data)
}
