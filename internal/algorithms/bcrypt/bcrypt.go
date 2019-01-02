package bcrypt

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

// Bcrypt implements password hashing algorithm interface
type Bcrypt struct{}

// DoHash hash given password string with bcrypt algorithm
func (Bcrypt) DoHash(pswd string) (pswdHash string) {
	byteHash, err := bcrypt.GenerateFromPassword([]byte(pswd), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	return string(byteHash)
}

// CheckHash compare matching with given password and hash with bcrypt algorithm
func (Bcrypt) CheckHash(pswd, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pswd))
	if err != nil {
		return false
	}
	return true
}
