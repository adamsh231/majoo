package repositories

import (
	"database/sql"
	"github.com/adamsh231/majoo/domain/interfaces"
	"github.com/adamsh231/majoo/domain/models"
)

type ProductRepository struct {
	PostgresDB *sql.DB
}

func NewProductRepository(postgresDB *sql.DB) interfaces.IProductRepository{
	return &ProductRepository{PostgresDB: postgresDB}
}

func (repo ProductRepository) Browse(search, orderBy, sort string, limit, offset int) (res []models.Product, err error) {
	panic("implement me")
}

func (repo ProductRepository) Read(model models.Product) (res models.Product, err error) {
	panic("implement me")
}

func (repo ProductRepository) Add(model models.Product, tx *sql.Tx) (res string, err error) {
	panic("implement me")
}

func (repo ProductRepository) Edit(model models.Product, tx *sql.Tx) (res string, err error) {
	panic("implement me")
}

func (repo ProductRepository) Delete(model models.Product, tx *sql.Tx) (res string, err error) {
	panic("implement me")
}

func (repo ProductRepository) Count(search string) (res int, err error) {
	panic("implement me")
}
