package algorithms

import (
	"testing"

	"github.com/andskur/pswd-hashing-tools/internal/algorithms/hash-passwords/argon2"
	"github.com/andskur/pswd-hashing-tools/internal/algorithms/hash-passwords/bcrypt"
	"github.com/andskur/pswd-hashing-tools/internal/algorithms/hash-passwords/scrypt"
)

var (
	testAlgos = Algorithms{
		Supported: map[string]HashAlgorithm{
			"bcrypt": &bcrypt.Bcrypt{},
			"scrypt": &scrypt.Scrypt{},
			"argon2": &argon2.Argon2{},
		},
		Default: "bcrypt",
	}

	trueAlgo  = testAlgos.RandomSupported()
	falseArgo = "damnhasher"
)

func TestSetAlgorithm(t *testing.T) {
	// Set algorithm type from name
	testAlgos.SetAlgorithm(trueAlgo)

	// Get name of received algorithm type
	algoName := GetName(testAlgos.Current)

	// Validate algo name
	if !testAlgos.ValidateAlgorithm(algoName) {
		t.Errorf("Given algorithm name %q is not valid", trueAlgo)
	}
}

func TestAlgorithms_ValidateAlgorithm(t *testing.T) {
	if !testAlgos.ValidateAlgorithm(trueAlgo) {
		t.Errorf("Function validated supported algorithm %q as unsupported", trueAlgo)
	}
	if testAlgos.ValidateAlgorithm(falseArgo) {
		t.Errorf("Function validated unsupported algorithm %q as supported", falseArgo)
	}
}

func TestRandomSupported(t *testing.T) {
	algo := testAlgos.RandomSupported()
	if !testAlgos.ValidateAlgorithm(trueAlgo) {
		t.Errorf("Function validated supported algorithm %q as unsupported", algo)
	}
	for keys := range testAlgos.Supported {
		if keys == algo {
			return
		}
	}
	t.Errorf("Function validated supported algorithm %q as unsupported", algo)
}
