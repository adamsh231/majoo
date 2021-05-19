package models

import (
	"database/sql"
	"time"
)

type Product struct {
	ID          string
	Merchant    Merchant
	Sku         string
	Name        string
	Slug        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   sql.NullTime
	Images      string
}

func NewProduct() *Product {
	return &Product{}
}

func (model Product) ScanRows(rows *sql.Rows) (res Product, err error) {
	err = rows.Scan(&res.ID, &res.Merchant.ID, &res.Merchant.Name, &res.Sku, &res.Name, &res.Slug, &res.Description, &res.CreatedAt, &res.UpdatedAt, &res.DeletedAt, &res.Images)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (model Product) ScanRow(row *sql.Row) (res Product, err error) {
	err = row.Scan(&res.ID, &res.Merchant.ID, &res.Merchant.Name, &res.Sku, &res.Name, &res.Slug, &res.Description, &res.CreatedAt, &res.UpdatedAt, &res.DeletedAt, &res.Images)
	if err != nil {
		return res, err
	}

	return res, nil
}
