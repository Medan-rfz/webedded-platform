package auth_service

import (
	"context"
	"time"

	auth_dto "users_management/internal/domain/dto/auth"
	"users_management/internal/domain/entities/errors"
	"users_management/internal/helpers/jwt"
	passwordhash "users_management/internal/helpers/password_hash"
)

func (s *authService) Login(ctx context.Context, data auth_dto.LoginDTO) (*auth_dto.LoginSuccess, error) {
	// TODO add tracing

	user, err := s.authRepo.GetAuthByEmail(ctx, data.Email)
	if err != nil {
		return nil, errors.ErrUserNotFound
	}

	if !passwordhash.CheckPassword(data.Password, user.PasswordHash) {
		return nil, errors.ErrInvalidPassword
	}

	token, err := jwt.CreateJwtToken(jwt.JwtData{
		Email:        data.Email,
		LiveTimeUnix: time.Now().Add(jwtLiveDuration).Unix(),
	}, jwtKey)
	if err != nil {
		return nil, err
	}

	refreshTimeLive := time.Now().Add(jwtRefreshLiveDuration)
	refreshToken, err := jwt.CreateJwtToken(jwt.JwtData{
		Email:        data.Email,
		LiveTimeUnix: refreshTimeLive.Unix(),
	}, jwtRefreshKey)
	if err != nil {
		return nil, err
	}

	s.authRepo.AddRefreshToken(ctx, auth_dto.AddRefreshDTO{
		UserId:       user.Id,
		RefreshToken: refreshToken,
		Expires:      refreshTimeLive,
	})

	return &auth_dto.LoginSuccess{
		Token:        token,
		RefreshToken: refreshToken,
	}, nil
}
