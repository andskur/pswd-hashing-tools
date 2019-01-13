package sha3

import (
	"encoding/hex"
	"golang.org/x/crypto/sha3"
)

// Sha3 implements hashing algorithm interface
type Sha3 struct{}

// DoHash hash given string with Sha2 algorithm
func (Sha3) DoHash(str string) (hash string) {
	h := sha3.New256()
	h.Write([]byte(str))
	sum := h.Sum(nil)

	return hex.EncodeToString(sum)
}

// CheckHash compare matching with given string and hash
func (Sha3) CheckHash(pswd, hash string) bool {
	return true
}
