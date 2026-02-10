package utils

import "golang.org/x/crypto/bcrypt"

type Bcrypt struct {
	Cost int
}

func (b *Bcrypt) HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), b.Cost)
}

func (c *Bcrypt) ComparePassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
