package auth_handler

import (
	"context"

	auth_dto "users_management/internal/domain/dto/auth"
)

func (h *authHandler) Refresh(ctx context.Context, data auth_dto.RefreshDTO) (string, error) {
	return h.authService.Refresh(ctx, data)
}
