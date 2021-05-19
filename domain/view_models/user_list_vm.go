package view_models

import (
	"github.com/adamsh231/majoo/domain/models"
)

type UserListVM struct {
	Name  string
	Email string
	//Role  string
}

func NewUserListVM() *UserListVM{
	return &UserListVM{}
}

func(vm *UserListVM) Build(model *models.User){
	*vm = UserListVM{
		Name:      model.Name,
		Email:     model.Email,
	}
}
