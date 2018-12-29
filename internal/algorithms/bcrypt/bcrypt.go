package bcrypt

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type Bcrypt struct{}

// DoHash hash given password string with bcrypt algorithm
func (Bcrypt) DoHash(pswd string) (pswdHash string) {
	byteHash, err := bcrypt.GenerateFromPassword([]byte(pswd), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	pswdHash = string(byteHash)

	return pswdHash
}

// CheckHash compare matching with given password and hash with bcrypt algorithm
func (Bcrypt) CheckHash(pswd, hash string) (result bool) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pswd))
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
