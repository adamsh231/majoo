package interfaces

import (
	"database/sql"
	"github.com/adamsh231/majoo/domain/models"
)

type IOutletRepository interface {
	Browse(search, orderBy, sort string, limit, offset int) (res []models.Outlet, err error)

	Read(model models.Outlet) (res models.Outlet, err error)

	Add(model models.Outlet, tx *sql.Tx) (res string, err error)

	Edit(model models.Outlet, tx *sql.Tx) (res string, err error)

	Delete(model models.Outlet, tx *sql.Tx) (res string, err error)

	Count(search string) (res int,err error)
}
