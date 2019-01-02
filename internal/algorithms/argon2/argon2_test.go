package argon2

import (
	"testing"
)

const (
	password = "qwerty123"
	hash     = "$argon2id$v=19$m=65536,t=3,p=2$d30dba4ef06d46b2706cf4753352c7b7$f52bc6a0cd02d90c29e24db31e8cf7bd67c761726a8679f86b0a4103c923d973"
)

var alg = &Argon2{}

func TestScrypt_CheckHash(t *testing.T) {
	if !alg.CheckHash(password, hash) {
		t.Errorf("True compare failed: %s not matching with %s", password, hash)
	}
}

func TestScrypt_DoHash(t *testing.T) {
	hash := alg.DoHash(password)
	err := CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		t.Errorf("Calculated hash not matching with given string")
	}
}
