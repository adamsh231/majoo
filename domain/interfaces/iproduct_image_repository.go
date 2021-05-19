package interfaces

import (
	"database/sql"
	"github.com/adamsh231/majoo/domain/models"
)

type IProductImageRepository interface {
	Browse(search, orderBy, sort string, limit, offset int) (res []models.ProductImage, err error)

	Read(model models.ProductImage) (res models.ProductImage, err error)

	Add(model models.ProductImage, tx *sql.Tx) (res string, err error)

	Edit(model models.ProductImage, tx *sql.Tx) (res string, err error)

	Delete(model models.ProductImage, tx *sql.Tx) (res string, err error)

	Count(search string) (res int,err error)
}
