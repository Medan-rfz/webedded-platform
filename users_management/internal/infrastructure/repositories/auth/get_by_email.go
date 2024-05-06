package auth_repo

import (
	"context"

	auth_entities "users_management/internal/domain/entities/auth"
	"users_management/internal/infrastructure/repositories"
)

func (r *authRepository) GetByEmail(
	ctx context.Context,
	email string,
	opts ...repositories.RepoOption,
) (*auth_entities.AuthUserModel, error) {
	q := r.getQuery(ctx, opts...)
	user, err := q.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return &auth_entities.AuthUserModel{
		Id:           user.ID,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
	}, nil
}
