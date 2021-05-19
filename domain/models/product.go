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
	Description sql.NullString
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   sql.NullTime
	Images      string
}

const(
	ProductSelectStatement = `SELECT P.id, M.id, M.name, P.sku, P.Name, P.Slug, P.description,
							  P.created_at, P.updated_at, P.deleted_at, array_to_string(array_agg(PI.path), ',') as images FROM products P 
							  INNER JOIN merchants M ON P.merchant_id = M.id
							  LEFT JOIN product_images PI ON PI.product_id = P.id`
	ProductDeleteStatement = `WHERE P.deleted_at IS NULL`
	ProductSearchStatement = `AND P.name LIKE $1`
	ProductGroupByStatement = `GROUP BY P.id, M.name, M.id`
)

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
