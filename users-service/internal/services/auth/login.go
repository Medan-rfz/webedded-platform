package auth_service

import (
	"context"
	"time"

	"webedded.users_management/internal/domain/entities/errors"
	"webedded.users_management/internal/helpers/jwt"

	auth_dto "webedded.users_management/internal/domain/dto/auth"

	passwordhash "webedded.users_management/internal/helpers/password_hash"
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
		UserId:       user.Id,
		Role:         user.Role,
		LiveTimeUnix: time.Now().Add(jwtLiveDuration).Unix(),
	}, jwtKey)
	if err != nil {
		return nil, err
	}

	refreshTimeLive := time.Now().Add(jwtRefreshLiveDuration)
	refreshToken, err := jwt.CreateJwtToken(jwt.JwtData{
		UserId:       user.Id,
		Role:         user.Role,
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
