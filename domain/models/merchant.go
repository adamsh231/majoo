package models

import "database/sql"

type Merchant struct {
	ID        string
	Name      string
	Phone     string
	Address   string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
}

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
