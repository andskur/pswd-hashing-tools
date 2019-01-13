package hash

//TODO implement FileCheckSum functions
//TODO implement Select key byte size

// PaswordHasher implements password hashing algorithms interface
type Hasher interface {
	// DoHash hash given password string
	DoHash(str string) (hash string)
	// CheckHash compare matching with given string and hash
	CheckHash(pswd, hash string) (result bool)
}
