package auth_entities

import "time"

type User struct {
	FirstName string
	LastName  string
	Email     string
	Group     string
}

type AuthUserModel struct {
	Id           int64
	Email        string
	PasswordHash string
}

type UserCredential struct {
	UserId       int64
	PasswordHash string
}

type UserRefreshToken struct {
	Id           int64
	RefreshToken string
	Expires      time.Time
}
