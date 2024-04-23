package auth_repo

import (
	"context"

	auth_entities "users_management/internal/domain/entities/auth"
	"users_management/internal/infrastructure/repositories"
)

func (r *authRepository) GetAuthById(
	ctx context.Context,
	id int64,
	opts ...repositories.RepoOption,
) (*auth_entities.AuthUserModel, error) {
	q := r.getQuery(opts...)

	row, err := q.GetAuthById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &auth_entities.AuthUserModel{
		Id:           row.ID,
		Email:        row.Email,
		PasswordHash: row.PasswordHash,
	}, nil
}
