package models

import (
	"database/sql"
	"time"
)

type User struct {
	ID        string
	Name      string
	Email     string
	Password  string
	Role      sql.NullString
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}

func NewUser() *User{
	return &User{}
}

const(
	UserSelectStatement = "SELECT U.id, U.name, U.email, U.password, U.role, U.created_at, U.updated_at, U.deleted_at FROM users U"
	UserDeleteStatement = "WHERE U.deleted_at IS NULL"
)

func (model User) ScanRows(rows *sql.Rows) (res User, err error) {
	err = rows.Scan(&res.ID, &res.Name, &res.Email, &res.Password, &res.Role, &res.CreatedAt, &res.UpdatedAt, &res.DeletedAt)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (model User) ScanRow(row *sql.Row) (res User, err error) {
	err = row.Scan(&res.ID, &res.Name, &res.Email, &res.Password, &res.Role, &res.CreatedAt, &res.UpdatedAt, &res.DeletedAt)
	if err != nil {
		return res, err
	}

	return res, nil
}