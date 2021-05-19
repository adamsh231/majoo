package repositories

import (
	"database/sql"
	"github.com/adamsh231/majoo/domain/interfaces"
	"github.com/adamsh231/majoo/domain/models"
	"strings"
)

type ProductRepository struct {
	PostgresDB *sql.DB
}

func NewProductRepository(postgresDB *sql.DB) interfaces.IProductRepository{
	return &ProductRepository{PostgresDB: postgresDB}
}

func (repo ProductRepository) Browse(search, orderBy, sort string, limit, offset int) (res []models.Product, err error) {
	model := models.NewProduct()
	statement := models.ProductSelectStatement + ` ` + models.ProductDeleteStatement + ` ` + models.ProductSearchStatement + ` ` + models.ProductGroupByStatement  + ` order by P.` + orderBy + ` ` + sort + ` limit $2 offset $3`
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

func (repo ProductRepository) Read(model models.Product) (res models.Product, err error) {
	statement := models.ProductSelectStatement + ` ` + models.ProductDeleteStatement + ` AND P.id=$1 ` + models.ProductGroupByStatement
	row := repo.PostgresDB.QueryRow(statement, model.ID)
	res, err = model.ScanRow(row)
	if err != nil {
		return res, err
	}

	return res, err
}

func (repo ProductRepository) Add(model models.Product, tx *sql.Tx) (res string, err error) {
	statement := `insert into products (merchant_id, sku, name, slug, description, created_at, updated_at) values ($1, $2, $3, $4, $5, $6, $7) returning id`
	err = tx.QueryRow(statement, model.Merchant.ID, model.Sku, model.Name, model.Slug, model.Description, model.CreatedAt, model.UpdatedAt).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, err
}

func (repo ProductRepository) Edit(model models.Product, tx *sql.Tx) (res string, err error) {
	statement := `UPDATE products SET merchant_id=$1, sku=$2, name=$3, slug=$4, description=$5, updated_at=$6 WHERE id=$7`
	_, err = tx.Exec(statement, model.Merchant.ID, model.Sku, model.Name, model.Slug, model.Description, model.UpdatedAt, model.ID)
	if err != nil {
		return res, err
	}

	res = model.ID
	return res, err
}

func (repo ProductRepository) Delete(model models.Product, tx *sql.Tx) (res string, err error) {
	statement := `UPDATE products SET deleted_at=$1 WHERE id=$2`
	_, err = tx.Exec(statement, model.DeletedAt, model.ID)
	if err != nil {
		return res, err
	}

	res = model.ID
	return res, err
}

func (repo ProductRepository) Count(search string) (res int, err error) {
	statement := `SELECT COUNT (P.id) FROM products P ` + models.ProductDeleteStatement + ` ` + models.ProductSearchStatement
	err = repo.PostgresDB.QueryRow(statement, "%"+strings.ToLower(search)+"%").Scan(&res)
	if err != nil {
		return res, err
	}

	return res, err
}
