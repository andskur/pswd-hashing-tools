package scrypt

import (
	"fmt"
	"github.com/elithrar/simple-scrypt"
	"log"
)

type Scrypt struct{}

func (Scrypt) DoHash(pswd string) (pswdHash string) {
	byteHash, err := scrypt.GenerateFromPassword([]byte(pswd), scrypt.DefaultParams)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(byteHash)
	pswdHash = string(byteHash)

	return pswdHash
}

func (Scrypt) CheckHash(pswd, hash string) (result bool) {
	fmt.Println([]byte(hash))
	err := scrypt.CompareHashAndPassword([]byte(hash), []byte(pswd))
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
