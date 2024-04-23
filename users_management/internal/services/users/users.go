package users_service

import (
	"context"
	"users_management/internal/domain/entities/users"

	authdto "users_management/internal/domain/dto/auth"
)

type usersService struct{}

func NewUsersService() *usersService {
	return &usersService{}
}

func (s *usersService) Create(ctx context.Context, user authdto.RegisterDTO) (*users.User, error) {
	// TODO Compute password hash

	return nil, nil
}

func (s *usersService) DeleteById(ctx context.Context, userId int64) error {
	return nil
}
