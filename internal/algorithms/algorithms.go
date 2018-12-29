package algorithms

// Algorithm implements password hashing algorithms interface
type Algorithm interface {
	// DoHash hash given password string
	DoHash(pswd string) (pswdHash string)

	// CheckHash compare matching with given password and hash
	CheckHash(pswd, hash string) (result bool)
}
