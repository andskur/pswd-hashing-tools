package bcrypt

import (
	"fmt"
	algo "golang.org/x/crypto/bcrypt"
	"log"
)

type Bcrypt struct{}

func (*Bcrypt) DoHash(pswd string) (pswdHash string, err error) {
	bytes, err := algo.GenerateFromPassword([]byte(pswd), algo.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	pswdHash = string(bytes)

	return
}

func (*Bcrypt) CheckHash(pswd, hash string) (result bool) {
	err := algo.CompareHashAndPassword([]byte(hash), []byte(pswd))
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}
