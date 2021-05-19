package interfaces

import (
	"github.com/adamsh231/majoo/domain/requests"
	"github.com/adamsh231/majoo/domain/view_models"
)

type IUserUseCase interface {
	Add(req *requests.UserAddRequest) (res string, err error)

	Edit(req *requests.UserEditRequest, id string) (res string, err error)

	Delete(id string) (res string, err error)

	Login(req *requests.UserLoginRequest) (res view_models.LoginVm, err error)
}
