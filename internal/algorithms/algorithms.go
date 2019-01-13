package algorithms

import (
	"fmt"
	"math/rand"
	"reflect"
	"strings"
	"time"
)

// Algorithms implements group of one typed crypto algorithms
type Algorithms struct {
	Supported map[string]HashAlgorithm
	Default   string
	Current   HashAlgorithm
}

// HashAlgorithm implements hashing algorithms interface
type HashAlgorithm interface {
	DoHash(string) string
	CheckHash(string, string) bool
}

// SetAlgorithm setting crypto hashing algorithm for interaction from given string
func (algos *Algorithms) SetAlgorithm(algoStr string) {
	if algos.ValidateAlgorithm(algoStr) {
		algos.Current = algos.Supported[algoStr]
	} else {
		fmt.Printf("%q algorithm doesn't supported, swith to default - %q\n", strings.Title(algoStr), strings.Title(algos.Default))
		algos.Current = algos.Supported[algos.Default]
	}
}

// ValidateAlgorithm check if the given algorithm is currently supported
func (algos *Algorithms) ValidateAlgorithm(check string) bool {
	for key := range algos.Supported {
		if key == check {
			return true
		}
	}
	return false
}

// RandomSupported get random supported crypto hashing algorithm
func (algos *Algorithms) RandomSupported() string {
	// Initialize global pseudo random generator
	rand.Seed(time.Now().Unix())

	// Ger array from supported algorithms map keys
	keys := reflect.ValueOf(algos.Supported).MapKeys()
	strKeys := make([]string, len(keys))
	for i := 0; i < len(keys); i++ {
		strKeys[i] = keys[i].String()
	}

	// Get random supported algorithm name
	return strKeys[rand.Intn(len(keys))]
}

// GetName get algorithm name from it type
func GetName(algo HashAlgorithm) string {
	algoType := reflect.TypeOf(algo).String()
	output := strings.Split(algoType, ".")
	algoName := strings.ToLower(output[1])

	return algoName
}
