package pw

type IService interface {
	HashAndEncrypt(password string) (*PasswordPair, error)
	Verify(password string, pair *PasswordPair) error
}
