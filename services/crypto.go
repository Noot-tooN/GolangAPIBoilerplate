package services

import "golang.org/x/crypto/bcrypt"

type ICryptoService interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
}

type CryptoService struct {}

func NewDefaultCryptoService() ICryptoService {
	return &CryptoService{}
}

func (cs CryptoService) HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func (cs CryptoService) CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}
