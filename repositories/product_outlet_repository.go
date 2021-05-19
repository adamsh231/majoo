package repositories

import (
	"database/sql"
	"github.com/adamsh231/majoo/domain/interfaces"
	"github.com/adamsh231/majoo/domain/models"
)

type ProductOutletRepository struct {
	PostgresDB *sql.DB
}

func NewProductOutletRepository(postgresDB *sql.DB) interfaces.IProductOutletRepository{
	return &ProductOutletRepository{PostgresDB: postgresDB}
}


func (repo ProductOutletRepository) Browse(search, orderBy, sort string, limit, offset int) (res []models.ProductOutlet, err error) {
	panic("implement me")
}

func (repo ProductOutletRepository) Read(model models.ProductOutlet) (res models.ProductOutlet, err error) {
	panic("implement me")
}

func (repo ProductOutletRepository) Add(model models.ProductOutlet, tx *sql.Tx) (res string, err error) {
	panic("implement me")
}

func (repo ProductOutletRepository) Edit(model models.ProductOutlet, tx *sql.Tx) (res string, err error) {
	panic("implement me")
}

func (repo ProductOutletRepository) Delete(model models.ProductOutlet, tx *sql.Tx) (res string, err error) {
	panic("implement me")
}

func (repo ProductOutletRepository) Count(search string) (res int, err error) {
	panic("implement me")
}
