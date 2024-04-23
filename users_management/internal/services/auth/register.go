package auth_service

import (
	"context"

	auth_dto "users_management/internal/domain/dto/auth"
	auth_entities "users_management/internal/domain/entities/auth"
	"users_management/internal/domain/entities/errors"
	passwordhash "users_management/internal/helpers/password_hash"
	"users_management/internal/infrastructure/repositories"

	"github.com/jackc/pgx/v5"
)

func (s *authService) Register(ctx context.Context, data auth_dto.RegisterDTO) error {
	// TODO add tracing

	user, err := s.authRepo.GetByEmail(ctx, data.Email)
	if err != nil && err != pgx.ErrNoRows {
		return err
	}

	if user != nil {
		return errors.ErrUserAlreadyExists
	}

	passwordHash, err := passwordhash.HashingPassword(data.Password)
	if err != nil {
		return err
	}

	tx, err := s.authRepo.NewTx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		tx.Rollback(ctx)
	}()

	userId, err := s.authRepo.CreateUser(ctx,
		auth_entities.User{
			FirstName: data.FirstName,
			LastName:  data.LastName,
			Email:     data.Email,
		},
		repositories.WithTx(tx))
	if err != nil {
		return err
	}

	s.authRepo.InsertCredential(ctx,
		auth_entities.UserCredential{
			UserId:       userId,
			PasswordHash: passwordHash,
		},
		repositories.WithTx(tx))

	return tx.Commit(ctx)
}
