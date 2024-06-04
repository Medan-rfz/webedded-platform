package jwt

import (
	"webedded.users_management/internal/domain/entities/errors"
	"webedded.users_management/internal/domain/entities/users"

	"github.com/golang-jwt/jwt/v5"
)

type JwtData struct {
	UserId       int64
	Role         users.Role
	LiveTimeUnix int64
}

func GetInfoFromJwtToken(token string, key []byte) (*JwtData, error) {
	parseToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := parseToken.Claims.(jwt.MapClaims)
	if !ok || !parseToken.Valid {
		return nil, errors.ErrParseToken
	}

	id, ok := claims["userId"].(int64)
	if !ok {
		return nil, errors.ErrParseToken
	}

	role, ok := claims["role"].(users.Role)
	if !ok {
		return nil, errors.ErrParseToken
	}

	liveTime, ok := claims["exp"].(int64)
	if !ok {
		return nil, errors.ErrParseToken
	}

	return &JwtData{
		UserId:       id,
		Role:         role,
		LiveTimeUnix: liveTime,
	}, nil
}

func CreateJwtToken(data JwtData, key []byte) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": data.UserId,
		"role":   data.Role,
		"exp":    data.LiveTimeUnix,
	})

	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
