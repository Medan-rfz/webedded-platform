package auth_repo

import (
	"context"

	auth_entities "webedded.users_management/internal/domain/entities/auth"
	"webedded.users_management/internal/infrastructure/repositories"
	users "webedded.users_management/internal/infrastructure/repositories/auth/queries"
)

func (r *authRepository) InsertCredential(
	ctx context.Context,
	user auth_entities.UserCredential,
	opts ...repositories.RepoOption,
) error {
	q := r.getQuery(ctx, opts...)
	return q.InsertCredentials(ctx, users.InsertCredentialsParams{
		UserID:       user.Id,
		PasswordHash: user.PasswordHash,
	})
}
