package argon2

import (
	"log"

	"github.com/andskur/argon2-hashing"
)

// Argon2 implements password hashing algorithm interface
type Argon2 struct{}

// DoHash hash given password string with argon2 algorithm
func (Argon2) DoHash(pswd string) (pswdHash string) {
	byteHash, err := argon2.GenerateFromPassword([]byte(pswd), argon2.DefaultParams)
	if err != nil {
		log.Fatal(err)
	}
	return string(byteHash)
}

// CheckHash compare matching with given password and hash with argon2 algorithm
func (Argon2) CheckHash(pswd, hash string) bool {
	err := argon2.CompareHashAndPassword([]byte(hash), []byte(pswd))
	if err != nil {
		return false
	}
	return true
}
