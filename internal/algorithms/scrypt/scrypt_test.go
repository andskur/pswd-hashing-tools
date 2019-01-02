package scrypt

import (
	"testing"

	"github.com/elithrar/simple-scrypt"
)

const (
	password = "qwerty123"
	hash     = "16384$8$1$9822477a20e30e070ecb7df6717413ec$2c537b14492b3265a139723022f32deea4f34c4d13b026ac98de369b0fcc5e66"
)

var alg = &Scrypt{}

func TestScrypt_CheckHash(t *testing.T) {
	if !alg.CheckHash(password, hash) {
		t.Errorf("True compare failed: %s not matching with %s", password, hash)
	}
}

func TestScrypt_DoHash(t *testing.T) {
	hash := alg.DoHash(password)
	err := scrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		t.Errorf("Calculated hash not matching with given string")
	}
}
