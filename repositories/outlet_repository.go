package repositories

import (
	"database/sql"
	"github.com/adamsh231/majoo/domain/interfaces"
	"github.com/adamsh231/majoo/domain/models"
	"strings"
)

type OutletRepository struct {
	PostgresDB *sql.DB
}

func NewOutletRepository(postgresDB *sql.DB) interfaces.IOutletRepository{
	return &OutletRepository{PostgresDB: postgresDB}
}

func (repo OutletRepository) Browse(search, orderBy, sort string, limit, offset int) (res []models.Outlet, err error) {
	model := models.NewOutlet()
	statement := models.OutletSelectStatement + ` ` + models.OutletDeleteStatement + ` ` + models.OutletSearchStatement + ` order by O.` + orderBy + ` ` + sort + ` limit $2 offset $3`
	rows, err := repo.PostgresDB.Query(statement, "%"+strings.ToLower(search)+"%", limit, offset)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		temp, err := model.ScanRows(rows)
		if err != nil {
			return res, err
		}
		res = append(res, temp)
	}

	return res, err
}

func (repo OutletRepository) Read(model models.Outlet) (res models.Outlet, err error) {
	panic("implement me")
}

func (repo OutletRepository) Add(model models.Outlet, tx *sql.Tx) (res string, err error) {
	statement := `insert into outlets (merchant_id, name, phone, address, created_at, updated_at) values ($1, $2, $3, $4, $5, $6) returning id`
	err = tx.QueryRow(statement, model.Merchant.ID, model.Name, model.Phone, model.Address, model.CreatedAt, model.UpdatedAt).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, err
}

func (repo OutletRepository) Edit(model models.Outlet, tx *sql.Tx) (res string, err error) {
	panic("implement me")
}

func (repo OutletRepository) Delete(model models.Outlet, tx *sql.Tx) (res string, err error) {
	panic("implement me")
}

func (repo OutletRepository) Count(search string) (res int, err error) {
	statement := `SELECT COUNT (O.id) FROM outlets O ` + models.OutletDeleteStatement + ` ` + models.OutletSearchStatement
	err = repo.PostgresDB.QueryRow(statement, "%"+strings.ToLower(search)+"%").Scan(&res)
	if err != nil {
		return res, err
	}

	return res, err
}