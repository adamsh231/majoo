package repositories

import (
	"database/sql"
	"github.com/adamsh231/majoo/domain/interfaces"
	"github.com/adamsh231/majoo/domain/models"
)

type MerchantRepository struct {
	PostgresDB *sql.DB
}

func NewMerchantRepository(postgresDB *sql.DB) interfaces.IMerchantRepository{
	return &MerchantRepository{PostgresDB: postgresDB}
}

func (repo MerchantRepository) Browse(search, orderBy, sort string, limit, offset int) (res []models.Merchant, err error) {
	panic("implement me")
}

func (repo MerchantRepository) Read(model models.Merchant) (res models.Merchant, err error) {
	panic("implement me")
}

func (repo MerchantRepository) Add(model models.Merchant, tx *sql.Tx) (res string, err error) {
	panic("implement me")
}

func (repo MerchantRepository) Edit(model models.Merchant, tx *sql.Tx) (res string, err error) {
	panic("implement me")
}

func (repo MerchantRepository) Delete(model models.Merchant, tx *sql.Tx) (res string, err error) {
	panic("implement me")
}

func (repo MerchantRepository) Count(search string) (res int, err error) {
	panic("implement me")
}
