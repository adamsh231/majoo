package interfaces

import (
	"database/sql"
	"github.com/adamsh231/majoo/domain/models"
)

type IProductOutletRepository interface {
	Browse(search, orderBy, sort string, limit, offset int) (res []models.ProductOutlet, err error)

	Read(model models.ProductOutlet) (res models.ProductOutlet, err error)

	Add(model models.ProductOutlet, tx *sql.Tx) (res string, err error)

	Edit(model models.ProductOutlet, tx *sql.Tx) (res string, err error)

	Delete(model models.ProductOutlet, tx *sql.Tx) (res string, err error)

	Count(search string) (res int,err error)
}
