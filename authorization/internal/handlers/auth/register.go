package auth_handler

import (
	"context"

	auth_dto "authorization/internal/domain/dto/auth"
	"authorization/internal/domain/entities/errors"
)

func (h *authHandler) Register(ctx context.Context, data auth_dto.RegisterDTO) error {
	isExist, err := h.authService.CheckUserExists(ctx, data)
	if err != nil {
		return err
	}

	if isExist {
		return errors.ErrUserAlreadyExists
	}

	createdUser, err := h.usersService.Create(ctx, data)
	if err != nil {
		return err
	}

	if err := h.authService.Register(ctx, auth_dto.AuthRegisterDTO{
		UserId:   createdUser.Id,
		Email:    data.Password,
		Password: data.Password,
	}); err != nil {
		h.usersService.DeleteById(ctx, createdUser.Id)
		return err
	}

	return nil
}
