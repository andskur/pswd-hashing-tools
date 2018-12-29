package scrypt

import (
	"fmt"
	"github.com/elithrar/simple-scrypt"
	"log"
)

type Scrypt struct{}

// DoHash hash given password string with scrypt algorithm
func (Scrypt) DoHash(pswd string) (pswdHash string) {
	byteHash, err := scrypt.GenerateFromPassword([]byte(pswd), scrypt.DefaultParams)
	if err != nil {
		log.Fatal(err)
	}
	pswdHash = string(byteHash)

	return pswdHash
}

// CheckHash compare matching with given password and hash with scrypt algorithm
func (Scrypt) CheckHash(pswd, hash string) (result bool) {
	err := scrypt.CompareHashAndPassword([]byte(hash), []byte(pswd))
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
