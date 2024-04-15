package auth_service

import (
	"context"
	"time"

	auth_dto "authorization/internal/domain/dto/auth"
	"authorization/internal/domain/entities/errors"
	passwordhash "authorization/internal/helpers/password_hash"

	"github.com/golang-jwt/jwt/v5"
)

func (s *authService) Login(ctx context.Context, data auth_dto.LoginDTO) (string, error) {
	user, err := s.usersRepo.GetByEmail(ctx, data.Email)
	if err != nil {
		return "", errors.ErrUserNotFound
	}

	if passwordhash.CheckPassword(data.Password, user.PasswordHash) {
		return "", errors.ErrInvalidPassword
	}

	return createJwtToken(data.Password)
}

func createJwtToken(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
