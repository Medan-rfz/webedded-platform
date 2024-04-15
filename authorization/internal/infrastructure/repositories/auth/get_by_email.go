package auth_repo

import (
	"context"

	auth_entities "authorization/internal/domain/entities/auth"
)

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*auth_entities.User, error) {
	user, err := r.queries.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return &auth_entities.User{
		Id:           user.UserID.Int64,
		Email:        user.Email,
		PasswordHash: user.PasswordHash.String,
	}, nil
}
