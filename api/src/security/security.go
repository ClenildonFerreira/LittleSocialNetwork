package security

import "golang.org/x/crypto/bcrypt"

// Hash receive a string and put a hash
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// CheckPassword compare a password and hash and returns if they match
func CheckPassword(passwordWithHash, passwordString string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwordWithHash), []byte(passwordString))
}
