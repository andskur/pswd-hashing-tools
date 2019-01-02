package argon2

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"strings"

	"golang.org/x/crypto/argon2"
)

//TODO need huge refactoring

// Argon2 implements password hashing algorithm interface
type Argon2 struct{}

// DoHash hash given password string with argon2 algorithm
func (Argon2) DoHash(pswd string) (pswdHash string) {
	byteHash, err := generateFromPassword([]byte(pswd), DefaultParams)
	if err != nil {
		log.Fatal(err)
	}
	return string(byteHash)
}

// CheckHash compare matching with given password and hash with argon2 algorithm
func (Argon2) CheckHash(pswd, hash string) (result bool) {
	err := comparePasswordAndHash(pswd, hash)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

//TODO move it to separate package

var (
	ErrInvalidHash         = errors.New("the encoded hash is not in the correct format")
	ErrIncompatibleVersion = errors.New("incompatible version of argon2")
	PasswordNotMatch       = errors.New("passwords not match")
)

type params struct {
	memory      uint32 // The amount of memory used by the algorithm (kibibytes)
	iterations  uint32 // The number of iterations (passes) over the memory.
	parallelism uint8  // The number of threads (lanes) used by the algorithm.
	saltLength  uint32 // Length of the random salt. 16 bytes is recommended for password hashing.
	keyLength   uint32 // Length of the generated key (password hash). 16 bytes or more is recommended.
}

// Default parameters
var DefaultParams = &params{
	memory:      64 * 1024,
	iterations:  3,
	parallelism: 2,
	saltLength:  16,
	keyLength:   32,
}

func generateFromPassword(password []byte, p *params) ([]byte, error) {
	// Generate a cryptographically secure random salt
	salt, err := generateRandomBytes(p.saltLength)
	if err != nil {
		return nil, err
	}

	// Pass the byte array password, salt and parameters to the argon2.IDKey
	// function. This will generate a hash of the password using the Argon2id
	// variant.
	key := argon2.IDKey(password, salt, p.iterations, p.memory, p.parallelism, p.keyLength)

	return []byte(fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%x$%x", argon2.Version, p.memory, p.iterations, p.parallelism, salt, key)), nil

	//b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	//b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	//encodedHash = fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, p.memory, p.iterations, p.parallelism, salt, key)

	//return encodedHash, nil
}

func generateRandomBytes(n uint32) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}

	return b, nil
}

func comparePasswordAndHash(password, encodedHash string) (err error) {
	p, salt, hash, err := decodeHash(encodedHash)
	if err != nil {
		log.Fatal(err)
	}

	otherHash := argon2.IDKey([]byte(password), salt, p.iterations, p.memory, p.parallelism, p.keyLength)

	if subtle.ConstantTimeCompare(hash, otherHash) != 1 {
		err = PasswordNotMatch
		return
	}
	return nil
}

func decodeHash(encodedHash string) (p *params, salt, hash []byte, err error) {
	vals := strings.Split(encodedHash, "$")
	if len(vals) != 6 {
		return nil, nil, nil, ErrInvalidHash
	}

	var version int
	_, err = fmt.Sscanf(vals[2], "v=%d", &version)
	if err != nil {
		return nil, nil, nil, err
	}
	if version != argon2.Version {
		return nil, nil, nil, ErrIncompatibleVersion
	}

	p = &params{}
	_, err = fmt.Sscanf(vals[3], "m=%d,t=%d,p=%d", &p.memory, &p.iterations, &p.parallelism)
	if err != nil {
		return nil, nil, nil, err
	}

	salt, err = base64.RawStdEncoding.DecodeString(vals[4])
	if err != nil {
		return nil, nil, nil, err
	}
	p.saltLength = uint32(len(salt))

	hash, err = base64.RawStdEncoding.DecodeString(vals[5])
	if err != nil {
		return nil, nil, nil, err
	}
	p.keyLength = uint32(len(hash))

	return p, salt, hash, nil
}
