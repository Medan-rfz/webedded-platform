package users_service

import (
	"context"

	authdto "authorization/internal/domain/dto/auth"
	"authorization/internal/domain/entities/users"
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
