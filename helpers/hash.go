package helpers

import "golang.org/x/crypto/bcrypt"

// bcrypt password
func HasPassword(password string) string {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashed)
}
