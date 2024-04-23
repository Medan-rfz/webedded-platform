package passwordhash

import "golang.org/x/crypto/bcrypt"

func HashingPassword(password string) (string, error) {
	pass, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(pass), err
}

func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}