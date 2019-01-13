package main

import (
	"github.com/andskur/pswd-hashing-tools/cmd"
	"github.com/andskur/pswd-hashing-tools/internal/algorithms"
	"github.com/andskur/pswd-hashing-tools/internal/algorithms/hash-passwords/argon2"
	"github.com/andskur/pswd-hashing-tools/internal/algorithms/hash-passwords/bcrypt"
	"github.com/andskur/pswd-hashing-tools/internal/algorithms/hash-passwords/scrypt"
)

// App entry point
func main() {
	pswdHashingAlgos := &algorithms.Algorithms{
		Supported: map[string]algorithms.HashAlgorithm{
			"bcrypt": &bcrypt.Bcrypt{},
			"scrypt": &scrypt.Scrypt{},
			"argon2": &argon2.Argon2{},
		},
		Default: "bcrypt",
	}

	cmd.Execute(pswdHashingAlgos)
}
