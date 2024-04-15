package auth_entities

type User struct {
	Id           int64
	Email        string
	PasswordHash string
}
