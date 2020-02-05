package account

import "golang.org/x/crypto/bcrypt"

//BCryptHasher implements PasswordHasher
type BCryptHasher struct {
}

//NewBCryptHasher return new instance of BCryptHasher
func NewBCryptHasher() *BCryptHasher {
	return &BCryptHasher{}
}

//Hash hash string (password)
func (me *BCryptHasher) Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 4)
	return string(bytes), err
}

//Verify verifies plain password and hashed password is correct
func (me *BCryptHasher) Verify(password string, hashedPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false, err
	}
	return true, nil
}
