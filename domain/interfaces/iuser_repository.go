package interfaces

import (
	"database/sql"
	"github.com/adamsh231/majoo/domain/models"
)

type IUserRepository interface {
	Read(model models.User) (res models.User, err error)

	Add(model models.User, tx *sql.Tx) (res string, err error)

	Edit(model models.User, tx *sql.Tx) (res string, err error)
}
