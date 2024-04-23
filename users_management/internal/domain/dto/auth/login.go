package auth_dto

type LoginDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginSuccess struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}
