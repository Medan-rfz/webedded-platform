package auth_repo

import (
	"context"

	auth_entities "webedded.users_management/internal/domain/entities/auth"
	"webedded.users_management/internal/infrastructure/repositories"
	users "webedded.users_management/internal/infrastructure/repositories/auth/queries"
)

func (r *authRepository) CreateUser(
	ctx context.Context,
	user auth_entities.User,
	opts ...repositories.RepoOption,
) (int64, error) {
	q := r.getQuery(ctx, opts...)
	return q.CreateUser(ctx, users.CreateUserParams{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	})
}
