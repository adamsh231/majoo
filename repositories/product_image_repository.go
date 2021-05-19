package repositories

import (
	"database/sql"
	"github.com/adamsh231/majoo/domain/interfaces"
	"github.com/adamsh231/majoo/domain/models"
)

type ProductImageRepository struct {
	PostgresDB *sql.DB
}

func NewProductImageRepository(postgresDB *sql.DB) interfaces.IProductImageRepository{
	return &ProductImageRepository{PostgresDB: postgresDB}
}

func (repo ProductImageRepository) Browse(search, orderBy, sort string, limit, offset int) (res []models.ProductImage, err error) {
	panic("implement me")
}

func (repo ProductImageRepository) Read(model models.ProductImage) (res models.ProductImage, err error) {
	panic("implement me")
}

func (repo ProductImageRepository) Add(model models.ProductImage, tx *sql.Tx) (res string, err error) {
	panic("implement me")
}

func (repo ProductImageRepository) Edit(model models.ProductImage, tx *sql.Tx) (res string, err error) {
	panic("implement me")
}

func (repo ProductImageRepository) Delete(model models.ProductImage, tx *sql.Tx) (res string, err error) {
	panic("implement me")
}

func (repo ProductImageRepository) Count(search string) (res int, err error) {
	panic("implement me")
}