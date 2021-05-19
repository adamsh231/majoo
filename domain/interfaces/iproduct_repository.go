package interfaces

import (
	"database/sql"
	"github.com/adamsh231/majoo/domain/models"
)

type IProductRepository interface {
	Browse(search, orderBy, sort string, limit, offset int) (res []models.Product, err error)

	Read(model models.Product) (res models.Product, err error)

	Add(model models.Product, tx *sql.Tx) (res string, err error)

	Edit(model models.Product, tx *sql.Tx) (res string, err error)

	Delete(model models.Product, tx *sql.Tx) (res string, err error)

	Count(search string) (res int,err error)
}
