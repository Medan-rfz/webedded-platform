package auth_service

import (
	"context"
	"time"

	auth_dto "webedded.users_management/internal/domain/dto/auth"
	"webedded.users_management/internal/domain/entities/errors"
	"webedded.users_management/internal/helpers/jwt"
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

	tokenInfo, err := jwt.GetInfoFromJwtToken(data.RefreshToken, jwtRefreshKey)
	if err != nil {
		return "", err
	}

	token, err := jwt.CreateJwtToken(jwt.JwtData{
		UserId:       tokenInfo.UserId,
		Role:         tokenInfo.Role,
		LiveTimeUnix: time.Now().Add(jwtLiveDuration).Unix(),
	}, jwtKey)
	if err != nil {
		return "", err
	}

	return token, nil
}
