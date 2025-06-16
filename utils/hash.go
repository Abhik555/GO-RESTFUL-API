package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	slc, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(slc), err
}

func CheckpasswordHash(pass, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))

	return err == nil
}
