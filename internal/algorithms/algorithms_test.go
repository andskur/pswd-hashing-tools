package algorithms

import (
	"reflect"
	"strings"
	"testing"
)

var (
	falseAlgo = "bestalgo"
	trueAlgo  = RandomSupported()
)

func TestSetAlgorithm(t *testing.T) {
	// Set algorithm type from name
	setAlgo, _ := SetAlgorithm(trueAlgo)

	// Get name of received algorithm type
	algoType := reflect.TypeOf(setAlgo).String()
	output := strings.Split(algoType, ".")
	algoName := strings.ToLower(output[1])

	// Validate algo name
	if !ValidateAlgorithm(algoName) {
		t.Errorf("Given algorithm name %q is not valid", trueAlgo)
	}
}

func TestValidateAlgorithm(t *testing.T) {
	if !ValidateAlgorithm(trueAlgo) {
		t.Errorf("Function validated supported algorithm %q as unsupported", trueAlgo)
	}
	if ValidateAlgorithm(falseAlgo) {
		t.Errorf("Function validated unsupported algorithm %q as supported", falseAlgo)
	}
}

func TestRandomSupported(t *testing.T) {
	algo := RandomSupported()
	if !ValidateAlgorithm(trueAlgo) {
		t.Errorf("Function validated supported algorithm %q as unsupported", algo)
	}
	for _, item := range Algos {
		if item == algo {
			return
		}
	}
	t.Errorf("Function validated supported algorithm %q as unsupported", algo)
}
