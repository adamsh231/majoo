package repositories

import (
	"database/sql"
	"github.com/adamsh231/majoo/domain/interfaces"
	"github.com/adamsh231/majoo/domain/models"
)

type OutletRepository struct {
	PostgresDB *sql.DB
}

func NewOutletRepository(postgresDB *sql.DB) interfaces.IOutletRepository{
	return &OutletRepository{PostgresDB: postgresDB}
}

func (repo OutletRepository) Browse(search, orderBy, sort string, limit, offset int) (res []models.Outlet, err error) {
	panic("implement me")
}

func (repo OutletRepository) Read(model models.Outlet) (res models.Outlet, err error) {
	panic("implement me")
}

func (repo OutletRepository) Add(model models.Outlet, tx *sql.Tx) (res string, err error) {
	panic("implement me")
}

func (repo OutletRepository) Edit(model models.Outlet, tx *sql.Tx) (res string, err error) {
	panic("implement me")
}

func (repo OutletRepository) Delete(model models.Outlet, tx *sql.Tx) (res string, err error) {
	panic("implement me")
}

func (repo OutletRepository) Count(search string) (res int, err error) {
	panic("implement me")
}