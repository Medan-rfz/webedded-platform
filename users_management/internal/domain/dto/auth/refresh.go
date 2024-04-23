package auth_dto

import "time"

type RefreshDTO struct {
	RefreshToken string `json:"refreshToken"`
}

type AddRefreshDTO struct {
	UserId       int64
	RefreshToken string
	Expires      time.Time
}
