package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

type HashHelper interface {
	Encrypt(str string) (string, error)
	CompareEncrypt(str, hash string) bool
}

type hashHelper struct {
}

type HHConfig struct {
}

func NewHashHelper(c *HHConfig) HashHelper {
	return &hashHelper{}
}

func (h *hashHelper) Encrypt(str string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func (h *hashHelper) CompareEncrypt(str, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(str))
	return err == nil
}
