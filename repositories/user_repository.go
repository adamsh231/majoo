package repositories

import (
	"database/sql"
	"github.com/adamsh231/majoo/domain/interfaces"
	"github.com/adamsh231/majoo/domain/models"
)

type UserRepository struct {
	PostgresDB *sql.DB
}

func NewUserRepository(postgresDB *sql.DB) interfaces.IUserRepository {
	return &UserRepository{PostgresDB: postgresDB}
}

func (repo UserRepository) Read(model models.User) (res models.User, err error) {
	statement := models.UserSelectStatement + ` ` + models.UserDeleteStatement + " AND U.email=$1"
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
