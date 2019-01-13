package main

import (
	"github.com/andskur/pswd-hashing-tools/cmd"
	"github.com/andskur/pswd-hashing-tools/internal/algorithms"
	"github.com/andskur/pswd-hashing-tools/internal/algorithms/hash-passwords/argon2"
	"github.com/andskur/pswd-hashing-tools/internal/algorithms/hash-passwords/bcrypt"
	"github.com/andskur/pswd-hashing-tools/internal/algorithms/hash-passwords/scrypt"
	"github.com/andskur/pswd-hashing-tools/internal/algorithms/hash/sha2"
	"github.com/andskur/pswd-hashing-tools/internal/algorithms/hash/sha3"
)

//TODO resolve interface problem

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

	hashingAlgos := &algorithms.Algorithms{
		Supported: map[string]algorithms.HashAlgorithm{
			"sha2": &sha2.Sha2{},
			"sha3": &sha3.Sha3{},
		},
		Default: "sha2",
	}

	cmd.Execute(pswdHashingAlgos, hashingAlgos)
}
