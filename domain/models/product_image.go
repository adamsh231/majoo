package models

import "database/sql"

type ProductImage struct {
	ID        string
	ProductID string
	Path      string
	Alt       string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
}

func NewProductImage() *ProductImage {
	return &ProductImage{}
}

func (model ProductImage) ScanRows(rows *sql.Rows) (res ProductImage, err error) {
	err = rows.Scan(&res.ID, &res.ProductID, &res.Path, &res.Alt, &res.CreatedAt, &res.UpdatedAt, &res.DeletedAt)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (model ProductImage) ScanRow(row *sql.Row) (res ProductImage, err error) {
	err = row.Scan(&res.ID, &res.ProductID, &res.Path, &res.Alt, &res.CreatedAt, &res.UpdatedAt, &res.DeletedAt)
	if err != nil {
		return res, err
	}

	return res, nil
}
