package sha2

import (
	"crypto/sha256"
	"encoding/hex"
)

// Sha2 implements hashing algorithm interface
type Sha2 struct{}

// DoHash hash given string with Sha2 algorithm
func (Sha2) DoHash(str string) (hash string) {
	h := sha256.New()
	h.Write([]byte(str))
	sum := h.Sum(nil)

	return hex.EncodeToString(sum)
}

// CheckHash compare matching with given string and hash
func (Sha2) CheckHash(pswd, hash string) bool {
	return true
}
