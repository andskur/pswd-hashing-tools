package scrypt

import (
	"log"

	"github.com/elithrar/simple-scrypt"
)

// Scrypt implements password hashing algorithm interface
type Scrypt struct{}

// DoHash hash given password string with scrypt algorithm
func (Scrypt) DoHash(pswd string) (pswdHash string) {
	byteHash, err := scrypt.GenerateFromPassword([]byte(pswd), scrypt.DefaultParams)
	if err != nil {
		log.Fatal(err)
	}
	return string(byteHash)
}

// CheckHash compare matching with given password and hash with scrypt algorithm
func (Scrypt) CheckHash(pswd, hash string) bool {
	err := scrypt.CompareHashAndPassword([]byte(hash), []byte(pswd))
	if err != nil {
		return false
	}
	return true
}
