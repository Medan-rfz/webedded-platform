package auth_service

import (
	"context"
	"time"

	auth_dto "users_management/internal/domain/dto/auth"
	"users_management/internal/domain/entities/errors"
	"users_management/internal/helpers/jwt"
)

func (s *authService) Refresh(ctx context.Context, data auth_dto.RefreshDTO) (string, error) {
	// Check exists refresh token in DB
	refreshTokenInfo, err := s.authRepo.GetRefreshToken(ctx, data.RefreshToken)
	if err != nil {
		return "", err
	}

	// Check time valid
	if refreshTokenInfo.Expires.Unix() < time.Now().Unix() {
		return "", errors.ErrRefreshTokenExpired
	}

	email, err := jwt.GetEmailFromJwtToken(data.RefreshToken, jwtRefreshKey)
	if err != nil {
		return "", err
	}

	token, err := jwt.CreateJwtToken(jwt.JwtData{
		Email:        email,
		LiveTimeUnix: time.Now().Add(jwtLiveDuration).Unix(),
	}, jwtKey)
	if err != nil {
		return "", err
	}

	return token, nil
}
