package hash

// PaswordHasher implements password hashing algorithms interface
type Hasher interface {
	// DoHash hash given password string
	DoHash(str string) (hash string)
}
