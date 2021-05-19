package view_models

import (
	"github.com/adamsh231/majoo/domain/models"
)

type UserListVM struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	//Role  string
}

func NewUserListVM() *UserListVM {
	return &UserListVM{}
}

func (vm *UserListVM) Build(model *models.User) {
	*vm = UserListVM{
		ID:    model.ID,
		Name:  model.Name,
		Email: model.Email,
	}
}
