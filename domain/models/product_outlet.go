package models

import "database/sql"

type ProductOutlet struct {
	ID        string
	Product   Product
	Outlet    Outlet
	Price     float64
	Stock     int
	CreatedAt string
	UpdatedAt string
	DeletedAt string
	Images    string
}

func NewProductOutlet() *ProductOutlet {
	return &ProductOutlet{}
}

func (model ProductOutlet) ScanRows(rows *sql.Rows) (res ProductOutlet, err error) {
	err = rows.Scan(&res.ID, &res.Product.ID, &res.Product.Sku, &res.Product.Name, &res.Outlet.ID, &res.Outlet.Name, &res.Price, &res.Stock, &res.CreatedAt, &res.UpdatedAt, &res.DeletedAt, &res.Images)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (model ProductOutlet) ScanRow(row *sql.Row) (res ProductOutlet, err error) {
	err = row.Scan(&res.ID, &res.Product.ID, &res.Product.Sku, &res.Product.Name, &res.Outlet.ID, &res.Outlet.Name, &res.Price, &res.Stock, &res.CreatedAt, &res.UpdatedAt, &res.DeletedAt, &res.Images)
	if err != nil {
		return res, err
	}

	return res, nil
}
