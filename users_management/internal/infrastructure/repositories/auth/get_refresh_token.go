package auth_repo

import (
	"context"

	auth_entities "users_management/internal/domain/entities/auth"
	"users_management/internal/infrastructure/repositories"
)

func (r *authRepository) GetRefreshToken(
	ctx context.Context,
	token string,
	opts ...repositories.RepoOption,
) (*auth_entities.UserRefreshToken, error) {
	q := r.getQuery(opts...)
	rTokenInfo, err := q.GetRefreshToken(ctx, token)
	if err != nil {
		return nil, err
	}

	return &auth_entities.UserRefreshToken{
		Id:           rTokenInfo.ID,
		RefreshToken: rTokenInfo.RefreshToken,
		Expires:      rTokenInfo.ExpiresAt.Time,
	}, nil
}
