package auth_dto

type RegisterDTO struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type AuthRegisterDTO struct {
	UserId   int64  `json:"userId"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
