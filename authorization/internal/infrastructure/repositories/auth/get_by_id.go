package auth_repo

import (
	"context"

	auth_entities "authorization/internal/domain/entities/auth"

	"github.com/jackc/pgx/v5/pgtype"
)

func (r *userRepository) GetById(ctx context.Context, id int64) (*auth_entities.User, error) {
	user, err := r.queries.GetById(ctx, pgtype.Int8{Int64: id, Valid: true})
	if err != nil {
		return nil, err
	}

	return &auth_entities.User{
		Id:           user.UserID.Int64,
		Email:        user.Email,
		PasswordHash: user.PasswordHash.String,
	}, nil
}
