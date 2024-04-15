package auth_repo

import (
	"context"

	auth_entities "authorization/internal/domain/entities/auth"
	users "authorization/internal/infrastructure/repositories/auth/queries"

	"github.com/jackc/pgx/v5/pgtype"
)

func (r *userRepository) Create(ctx context.Context, user auth_entities.User) error {
	return r.queries.Create(ctx, users.CreateParams{
		UserID:       pgtype.Int8{Int64: user.Id, Valid: true},
		Email:        user.Email,
		PasswordHash: pgtype.Text{String: user.PasswordHash, Valid: true},
	})
}
