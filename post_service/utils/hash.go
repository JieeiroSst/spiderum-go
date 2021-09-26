package utils

import (
	"golang.org/x/crypto/bcrypt"
)

type Hash struct {}

func NewHash()*Hash{
	return &Hash{}
}

func (h *Hash) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func (h *Hash) CheckPassowrd(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err
}