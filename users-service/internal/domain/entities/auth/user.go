package auth_entities

import (
	"time"

	"webedded.users_management/internal/domain/entities/users"
)

type User struct {
	FirstName string
	LastName  string
	Email     string
	Group     string
}

type AuthUserModel struct {
	Id           int64
	Role         users.Role
	Email        string
	PasswordHash string
}

type UserCredential struct {
	Id           int64
	PasswordHash string
}

type UserRefreshToken struct {
	Id           int64
	RefreshToken string
	Expires      time.Time
}
