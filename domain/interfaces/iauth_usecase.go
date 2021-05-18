package interfaces

type IAuthUseCase interface {
	Login(email, password string) (res string, err error)
}