package bcrypt

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

const (
	password = "qwerty123"
	hash     = "$2a$10$1eMGEMyY6xs0BbcukkBFVeiEmOWLW9LaT1CZihCqK384V3OJkyA9K"
)

var alg = &Bcrypt{}

func TestScrypt_CheckHash(t *testing.T) {
	if !alg.CheckHash(password, hash) {
		t.Errorf("True compare failed: %s not matching with %s", password, hash)
	}
}

func TestScrypt_DoHash(t *testing.T) {
	hash := alg.DoHash(password)
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		t.Errorf("Calculated hash not matching with given string")
	}
}
