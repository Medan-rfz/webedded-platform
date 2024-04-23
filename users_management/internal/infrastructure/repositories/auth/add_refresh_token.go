package auth_repo

import (
	"context"

	auth_dto "users_management/internal/domain/dto/auth"
	"users_management/internal/infrastructure/repositories"
	auth "users_management/internal/infrastructure/repositories/auth/queries"

	"github.com/jackc/pgx/v5/pgtype"
)

func (r *authRepository) AddRefreshToken(
	ctx context.Context,
	token auth_dto.AddRefreshDTO,
	opts ...repositories.RepoOption,
) error {
	q := r.getQuery(opts...)
	return q.AddRefreshToken(ctx, auth.AddRefreshTokenParams{
		UserID:       pgtype.Int8{Int64: token.UserId, Valid: true},
		RefreshToken: token.RefreshToken,
		ExpiresAt:    pgtype.Timestamp{Time: token.Expires, Valid: true},
	})
}
