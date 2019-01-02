package argon2

import (
	"testing"

	"github.com/andskur/argon2-hashing"
)

const (
	password = "qwerty123"
	hash     = "argon2id$19$65536$3$2$T8wT/pIJDqo/UrVHXXS8Ag$vaC5FaJRoVAnkxXcQCVnmeJ4lLp5Mp3NKRqDfXDcVv8"
)

var alg = &Argon2{}

func TestScrypt_CheckHash(t *testing.T) {
	if !alg.CheckHash(password, hash) {
		t.Errorf("True compare failed: %s not matching with %s", password, hash)
	}
}

func TestScrypt_DoHash(t *testing.T) {
	hash := alg.DoHash(password)
	err := argon2.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		t.Errorf("Calculated hash not matching with given string")
	}
}
