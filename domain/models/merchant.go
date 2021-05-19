package models

import (
	"database/sql"
	"time"
)

type Merchant struct {
	ID        string
	Name      string
	Phone     string
	Address   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}

const(
	MerchantSelectStatement = `SELECT M.id, M.name, M.phone, M.address, M.created_at, M.updated_at, M.deleted_at FROM merchants M`
	MerchantDeleteStatement = `WHERE M.deleted_at IS NULL`
	MerchantSearchStatement = `AND M.name LIKE $1`
)

func NewMerchant() *Merchant {
	return &Merchant{}
}

func (model Merchant) ScanRows(rows *sql.Rows) (res Merchant, err error) {
	err = rows.Scan(&res.ID, &res.Name, &res.Phone, &res.Address, &res.CreatedAt, &res.UpdatedAt, &res.DeletedAt)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (model Merchant) ScanRow(row *sql.Row) (res Merchant, err error) {
	err = row.Scan(&res.ID, &res.Name, &res.Phone, &res.Address, &res.CreatedAt, &res.UpdatedAt, &res.DeletedAt)
	if err != nil {
		return res, err
	}

	return res, nil
}
