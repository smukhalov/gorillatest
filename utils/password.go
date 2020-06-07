package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) ([]byte, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return bytes, err
}

func CheckPasswordHash(x, y []byte) bool {
	if len(x) != len(y) {
		return false
	}

	for i := 0; i < len(x); i++ {
		if x[i]^y[i] == 1 {
			return false
		}
	}
	return true
	//return subtle.ConstantTimeCompare(password, hash) == 1
	// err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	// if err == nil {
	// 	return nil
	// }

	// return err
}
