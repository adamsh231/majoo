package usecase

import (
	"github.com/adamsh231/majoo/domain/interfaces"
	"github.com/adamsh231/majoo/packages/helper"
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

	isValid := helper.CheckHashString(password, credPass)
	if isValid{
		res = "Logged In"
	}

	return res, err
}

