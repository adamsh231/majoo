package repositories

import (
	"database/sql"
	"github.com/adamsh231/majoo/domain/interfaces"
	"github.com/adamsh231/majoo/domain/models"
	"strings"
)

type MerchantRepository struct {
	PostgresDB *sql.DB
}

func NewMerchantRepository(postgresDB *sql.DB) interfaces.IMerchantRepository{
	return &MerchantRepository{PostgresDB: postgresDB}
}

func (repo MerchantRepository) Browse(search, orderBy, sort string, limit, offset int) (res []models.Merchant, err error) {
	model := models.NewMerchant()
	statement := models.MerchantSelectStatement + ` ` + models.MerchantDeleteStatement + ` ` + models.MerchantSearchStatement + ` order by M.` + orderBy + ` ` + sort + ` limit $2 offset $3`
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

func (repo MerchantRepository) Read(model models.Merchant) (res models.Merchant, err error) {
	statement := models.MerchantSelectStatement + ` ` + models.MerchantDeleteStatement + ` AND M.id=$1 `
	row := repo.PostgresDB.QueryRow(statement, model.ID)
	res, err = model.ScanRow(row)
	if err != nil {
		return res, err
	}

	return res, err
}

func (repo MerchantRepository) Add(model models.Merchant, tx *sql.Tx) (res string, err error) {
	statement := `insert into merchants (name, phone, address, created_at, updated_at) values ($1, $2, $3, $4, $5) returning id`
	err = tx.QueryRow(statement, model.Name, model.Phone, model.Address, model.CreatedAt, model.UpdatedAt).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, err
}

func (repo MerchantRepository) Edit(model models.Merchant, tx *sql.Tx) (res string, err error) {
	panic("implement me")
}

func (repo MerchantRepository) Delete(model models.Merchant, tx *sql.Tx) (res string, err error) {
	panic("implement me")
}

func (repo MerchantRepository) Count(search string) (res int, err error) {
	statement := `SELECT COUNT (M.id) FROM merchants M ` + models.MerchantDeleteStatement + ` ` + models.MerchantSearchStatement
	err = repo.PostgresDB.QueryRow(statement, "%"+strings.ToLower(search)+"%").Scan(&res)
	if err != nil {
		return res, err
	}

	return res, err
}
