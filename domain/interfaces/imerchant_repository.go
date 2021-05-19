package interfaces

import (
	"database/sql"
	"github.com/adamsh231/majoo/domain/models"
)

type IMerchantRepository interface {
	Browse(search, orderBy, sort string, limit, offset int) (res []models.Merchant, err error)

	Read(model models.Merchant) (res models.Merchant, err error)

	Add(model models.Merchant, tx *sql.Tx) (res string, err error)

	Edit(model models.Merchant, tx *sql.Tx) (res string, err error)

	Delete(model models.Merchant, tx *sql.Tx) (res string, err error)

	Count(search string) (res int,err error)
}
