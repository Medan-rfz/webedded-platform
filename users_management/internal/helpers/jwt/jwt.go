package jwt

import (
	"users_management/internal/domain/entities/errors"

	"github.com/golang-jwt/jwt/v5"
)

type JwtData struct {
	Email        string
	LiveTimeUnix int64
}

func GetEmailFromJwtToken(token string, key []byte) (string, error) {
	parseToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := parseToken.Claims.(jwt.MapClaims)
	if !ok || !parseToken.Valid {
		return "", errors.ErrParseToken
	}

	email, ok := claims["email"].(string)
	if !ok {
		return "", errors.ErrParseToken
	}

	return email, nil
}

func CreateJwtToken(data JwtData, key []byte) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": data.Email,
		"exp":   data.LiveTimeUnix,
	})

	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
