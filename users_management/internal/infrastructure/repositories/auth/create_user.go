package auth_repo

import (
	"context"

	auth_entities "users_management/internal/domain/entities/auth"
	"users_management/internal/infrastructure/repositories"
	users "users_management/internal/infrastructure/repositories/auth/queries"
)

func (r *authRepository) CreateUser(
	ctx context.Context,
	user auth_entities.User,
	opts ...repositories.RepoOption,
) (int64, error) {
	q := r.getQuery(opts...)
	return q.CreateUser(ctx, users.CreateUserParams{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	})
}
