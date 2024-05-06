package auth_repo

import (
	"context"

	auth_entities "users_management/internal/domain/entities/auth"
	"users_management/internal/infrastructure/repositories"
	users "users_management/internal/infrastructure/repositories/auth/queries"
)

func (r *authRepository) InsertCredential(
	ctx context.Context,
	user auth_entities.UserCredential,
	opts ...repositories.RepoOption,
) error {
	q := r.getQuery(ctx, opts...)
	return q.InsertCredentials(ctx, users.InsertCredentialsParams{
		UserID:       user.UserId,
		PasswordHash: user.PasswordHash,
	})
}
