package algorithms

// Algorithm implements password hashing algorithms interface
type Algorithm interface {
	DoHash(pswd string) (pswdHash string, err error)
	CheckHash(pswd, hash string) (result bool)
}
