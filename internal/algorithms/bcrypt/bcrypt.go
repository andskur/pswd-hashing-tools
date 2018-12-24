package bcrypt

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type Bcrypt struct{}

func (Bcrypt) DoHash(pswd string) (pswdHash string, err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pswd), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	pswdHash = string(bytes)

	return pswdHash, nil
}

func (Bcrypt) CheckHash(pswd, hash string) (result bool) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pswd))
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
