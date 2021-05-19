package models

import "database/sql"

type Outlet struct {
	ID         string
	MerchantID string
	Name       string
	Phone      string
	Address    string
	CreatedAt  string
	UpdatedAt  string
	DeletedAt  string
}

func NewOutlet() *Outlet {
	return &Outlet{}
}

func (model Outlet) ScanRows(rows *sql.Rows) (res Outlet, err error) {
	err = rows.Scan(&res.ID, &res.MerchantID, &res.Name, &res.Phone, &res.Address, &res.CreatedAt, &res.UpdatedAt, &res.DeletedAt)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (model Outlet) ScanRow(row *sql.Row) (res Outlet, err error) {
	err = row.Scan(&res.ID, &res.MerchantID, &res.Name, &res.Phone, &res.Address, &res.CreatedAt, &res.UpdatedAt, &res.DeletedAt)
	if err != nil {
		return res, err
	}

	return res, nil
}