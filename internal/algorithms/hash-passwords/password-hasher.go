package hash_passwords

// PaswordHasher implements password hashing algorithms interface
type PaswordHasher interface {
	// DoHash hash given password string
	DoHash(pswd string) (pswdHash string)

	// CheckHash compare matching with given password and hash
	CheckHash(pswd, hash string) (result bool)
}
