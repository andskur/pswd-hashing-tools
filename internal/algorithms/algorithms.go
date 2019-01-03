package algorithms

import (
	"fmt"
	"strings"

	"github.com/andskur/pswd-hashing-tools/internal/algorithms/argon2"
	"github.com/andskur/pswd-hashing-tools/internal/algorithms/bcrypt"
	"github.com/andskur/pswd-hashing-tools/internal/algorithms/scrypt"
)

// Algorithm implements password hashing algorithms interface
type Algorithm interface {
	// DoHash hash given password string
	DoHash(pswd string) (pswdHash string)

	// CheckHash compare matching with given password and hash
	CheckHash(pswd, hash string) (result bool)
}

var (
	algos       = []string{"bcrypt", "scrypt", "argon2"} // Suppported crypto algorithms
	defaultAlgo = "bcrypt"                               // Default crypto algorithm
)

// SetAlgorithm setting crypto hashing algorithm for interaction from given string
func SetAlgorithm(algoStr string) (algo Algorithm) {
	switch algoStr {
	case "argon2":
		algo = argon2.Argon2{}
	case "bcrypt":
		algo = &bcrypt.Bcrypt{}
	case "scrypt":
		algo = &scrypt.Scrypt{}
	default:
		// Instead of a separate validation function
		fmt.Printf("%q algorithm doesn't supported, swith to default - %q\n", strings.Title(algoStr), strings.Title(defaultAlgo))
		algoStr = defaultAlgo
		algo = &bcrypt.Bcrypt{}
	}
	fmt.Printf("Using %q hashing algorithm \n", strings.Title(algoStr))
	return algo
}

// ValidateAlgorithm check if the given algorithm is currently supported
func ValidateAlgorithm(check string) bool {
	fmt.Println(algos)
	for _, item := range algos {
		if item == check {
			return true
		}
	}
	return false
}
