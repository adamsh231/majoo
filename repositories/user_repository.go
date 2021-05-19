package repositories

import (
	"database/sql"
	"github.com/adamsh231/majoo/domain/interfaces"
	"github.com/adamsh231/majoo/domain/models"
	"strings"
)

type UserRepository struct {
	PostgresDB *sql.DB
}

func NewUserRepository(postgresDB *sql.DB) interfaces.IUserRepository {
	return &UserRepository{PostgresDB: postgresDB}
}

func (repo UserRepository) Browse(search, orderBy, sort string, limit, offset int) (res []models.User, err error) {
	model := models.NewUser()
	statement := models.UserSelectStatement + ` ` + models.UserDeleteStatement + ` ` + models.UserSearchStatement + ` order by U.` + orderBy + ` ` + sort + ` limit $2 offset $3`
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

	return res, nil
}

func (repo UserRepository) Read(model models.User) (res models.User, err error) {
	statement := models.UserSelectStatement + ` ` + models.UserDeleteStatement + ` ` + models.UserSearchStatement
	row := repo.PostgresDB.QueryRow(statement, model.Email)
	res, err = model.ScanRow(row)
	if err != nil {
		return res, err
	}

	return res, err
}

func (repo UserRepository) Add(model models.User, tx *sql.Tx) (res string, err error) {
	statement := `insert into users (name, email, password, created_at, updated_at) values ($1, $2, $3, $4, $5) returning id`
	err = tx.QueryRow(statement, model.Name, model.Email, model.Password, model.CreatedAt, model.UpdatedAt).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, err
}

func (repo UserRepository) Edit(model models.User, tx *sql.Tx) (res string, err error) {
	statement := `UPDATE users SET name=$1, email=$2, password=$3, updated_at=$4 WHERE id=$5`
	_, err = tx.Exec(statement, model.Name, model.Email, model.Password, model.UpdatedAt, model.ID)
	if err != nil {
		return res, err
	}

	res = model.ID
	return res, err
}

func (repo UserRepository) Delete(model models.User, tx *sql.Tx) (res string, err error) {
	statement := `UPDATE users SET deleted_at=$1 WHERE id=$2`
	_, err = tx.Exec(statement, model.DeletedAt, model.ID)
	if err != nil {
		return res, err
	}

	res = model.ID
	return res, err
}

func (repo UserRepository) Count(search string) (res int, err error) {
	statement := `SELECT COUNT (U.id) FROM users U ` + models.UserDeleteStatement + ` ` + models.UserSearchStatement
	err = repo.PostgresDB.QueryRow(statement, "%"+strings.ToLower(search)+"%").Scan(&res)
	if err != nil {
		return res, err
	}

	return res, err
}