package auth_service

import (
	"context"

	auth_dto "authorization/internal/domain/dto/auth"
)

func (s *authService) CheckUserExists(ctx context.Context, data auth_dto.RegisterDTO) (bool, error) {
	user, err := s.usersRepo.GetByEmail(ctx, data.Email)
	if err != nil {
		return false, err
	}

	if user != nil {
		return true, nil
	}

	return false, nil
}
