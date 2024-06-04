package auth_service

import (
	"context"

	auth_dto "webedded.users_management/internal/domain/dto/auth"
	auth_entities "webedded.users_management/internal/domain/entities/auth"
	"webedded.users_management/internal/domain/entities/errors"
	passwordhash "webedded.users_management/internal/helpers/password_hash"

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

	err = s.txExecutor.TxBegin(ctx, func(ctx context.Context) error {
		userId, err := s.authRepo.CreateUser(ctx,
			auth_entities.User{
				FirstName: data.FirstName,
				LastName:  data.LastName,
				Email:     data.Email,
			})
		if err != nil {
			return err
		}

		return s.authRepo.InsertCredential(ctx,
			auth_entities.UserCredential{
				Id:           userId,
				PasswordHash: passwordHash,
			})
	})

	return err
}
