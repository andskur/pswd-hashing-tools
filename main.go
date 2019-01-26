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

//TODO add viper package for configuration

// App entry point
func main() {
	app := initApp()

	cmd.Execute(app.PswdHashingAlgos, app.PswdHashingAlgos)
}

// Application structure implement application common parameters
type Application struct {
	HashingAlgos     *algorithms.Algorithms
	PswdHashingAlgos *algorithms.Algorithms
}

// Init Application parameters
func initApp() Application {
	return Application{
		PswdHashingAlgos: &algorithms.Algorithms{
			Supported: map[string]algorithms.HashAlgorithm{
				"bcrypt": &bcrypt.Bcrypt{},
				"scrypt": &scrypt.Scrypt{},
				"argon2": &argon2.Argon2{},
			},
			Default: "bcrypt",
		},
		HashingAlgos: &algorithms.Algorithms{
			Supported: map[string]algorithms.HashAlgorithm{
				"sha2": &sha2.Sha2{},
				"sha3": &sha3.Sha3{},
			},
			Default: "sha2",
		},
	}
}
