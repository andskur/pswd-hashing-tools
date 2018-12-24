package algorithms

// Algorithm implements password hashing algorithms interface
type Algorithm interface {
	DoHash(pswd string) (pswdHash string)
	CheckHash(pswd, hash string) (result bool)
}
