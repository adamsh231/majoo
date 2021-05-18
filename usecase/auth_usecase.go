package usecase

import (
	"github.com/adamsh231/majoo/domain/interfaces"
	"golang.org/x/crypto/bcrypt"
)

type AuthUseCase struct {
	*Contract
}

func NewAuthUseCase(contract *Contract) interfaces.IAuthUseCase {
	return &AuthUseCase{Contract: contract}
}

func (a AuthUseCase) Login(email, password string) (res string, err error) {
	credEmail := "adamsyarif219@gmail.com"
	credPass := "$2a$04$o4DYkkn9Frze9eXT5OfpteYvvzvSeOPLXhHaTJSkX/Uq2TzT2JotG"

	if email != credEmail{
		return "Email salah!", err
	}

	isValid := checkHashString(password, credPass)
	if isValid{
		res = "Password benar"
	}

	return res, err
}

func checkHashString(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}