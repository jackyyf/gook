package password

// Using bcrypt as password hasher.

import (
	"golang.org/x/crypto/bcrypt"
)

const difficulty = 8

func Generate(password string) string {
	r, err := bcrypt.GenerateFromPassword([]byte(password), difficulty)
	if err != nil {
		// Unrecoverable, panic to prevent store garbage password.
		panic(err)
	}
	return string(r)
}

func Check(password, hashed string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password)) == nil
}
