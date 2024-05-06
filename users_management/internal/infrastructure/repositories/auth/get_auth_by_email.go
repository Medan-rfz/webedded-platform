package auth_repo

import (
	"context"

	auth_entities "users_management/internal/domain/entities/auth"
	"users_management/internal/infrastructure/repositories"
)

func (r *authRepository) GetAuthByEmail(
	ctx context.Context,
	email string,
	opts ...repositories.RepoOption,
) (*auth_entities.AuthUserModel, error) {
	q := r.getQuery(ctx, opts...)

	row, err := q.GetAuthByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return &auth_entities.AuthUserModel{
		Id:           row.ID,
		Email:        row.Email,
		PasswordHash: row.PasswordHash,
	}, nil
}
