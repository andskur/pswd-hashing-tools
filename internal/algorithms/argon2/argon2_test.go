package argon2

import (
	"testing"
)

const (
	password = "qwerty123"
	hash     = "$argon2id$v=19$m=65536,t=3,p=2$0NHM2VjdVMWeVg0xgaNqzw$CnKdmqdOoeIT83alh1wQTEVzRIvuJ9iqVQVMQ2nwzZE"
)

var alg = &Argon2{}

func TestScrypt_CheckHash(t *testing.T) {
	if !alg.CheckHash(password, hash) {
		t.Errorf("True compare failed: %s not matching with %s", password, hash)
	}
}

func TestScrypt_DoHash(t *testing.T) {
	hash := alg.DoHash(password)
	err := comparePasswordAndHash(password, hash)
	if err != nil {
		t.Errorf("Calculated hash not matching with given string")
	}
}
