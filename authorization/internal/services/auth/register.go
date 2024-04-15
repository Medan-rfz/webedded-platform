package auth_service

import (
	"context"

	auth_dto "authorization/internal/domain/dto/auth"
	auth_entities "authorization/internal/domain/entities/auth"
	"authorization/internal/domain/entities/errors"
	passwordhash "authorization/internal/helpers/password_hash"
)

func (s *authService) Register(ctx context.Context, data auth_dto.AuthRegisterDTO) error {
	user, err := s.usersRepo.GetByEmail(ctx, data.Email)
	if err != nil {
		return err
	}

	if user != nil {
		return errors.ErrUserAlreadyExists
	}

	passwordHash, err := passwordhash.HashingPassword(data.Password)
	if err != nil {
		return err
	}

	return s.usersRepo.Create(ctx, auth_entities.User{
		Id:           data.UserId,
		Email:        data.Email,
		PasswordHash: passwordHash,
	})
}
