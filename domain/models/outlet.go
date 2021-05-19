package models

import (
	"database/sql"
	"time"
)

type Outlet struct {
	ID        string
	Merchant  Merchant
	Name      string
	Phone     string
	Address   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}

const(
	OutletSelectStatement = `SELECT O.id, O.name, O.phone, O.address, O.created_at, O.updated_at, O.deleted_at FROM outlets O`
	OutletDeleteStatement = `WHERE O.deleted_at IS NULL`
	OutletSearchStatement = `AND O.name LIKE $1`
)


func NewOutlet() *Outlet {
	return &Outlet{}
}

func (model Outlet) ScanRows(rows *sql.Rows) (res Outlet, err error) {
	err = rows.Scan(&res.ID, &res.Name, &res.Phone, &res.Address, &res.CreatedAt, &res.UpdatedAt, &res.DeletedAt)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (model Outlet) ScanRow(row *sql.Row) (res Outlet, err error) {
	err = row.Scan(&res.ID, &res.Name, &res.Phone, &res.Address, &res.CreatedAt, &res.UpdatedAt, &res.DeletedAt)
	if err != nil {
		return res, err
	}

	return res, nil
}
