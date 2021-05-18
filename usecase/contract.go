package usecase

import (
	"database/sql"
	"github.com/adamsh231/majoo/domain/view_models"
	"github.com/adamsh231/majoo/packages/jwe"
	"github.com/adamsh231/majoo/packages/jwt"
	ut "github.com/go-playground/universal-translator"
	validator "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Contract struct {
	UserID        string
	RoleID        int
	App           *fiber.App
	PostgresDB    *sql.DB
	PostgresTX    *sql.Tx
	Validate      *validator.Validate
	Translator    ut.Translator
	JweCredential jwe.Credential
	JwtCredential jwt.JwtCredential
}

const (
	// default limit for pagination
	defaultLimit = 10

	// max limit for pagination
	maxLimit = 50

	// default order by
	defaultOrderBy = "id"

	// default sort
	defaultSort = "asc"

	// default last page for pagination
	defaultLastPage = 0
)

func (uc Contract) setPaginationParameter(page, limit int, order, sort string) (int, int, int, string, string) {
	if page <= 0 {
		page = 1
	}
	if limit <= 0 || limit > maxLimit {
		limit = defaultLimit
	}
	if order == "" {
		order = defaultOrderBy
	}
	if sort == "" {
		sort = defaultSort
	}
	offset := (page - 1) * limit

	return offset, limit, page, order, sort
}

func (uc Contract) setPaginationResponse(page, limit, total int) (res view_models.PaginationVm) {
	var lastPage int

	if total > 0 {
		lastPage = total / limit

		if total%limit != 0 {
			lastPage = lastPage + 1
		}
	} else {
		lastPage = defaultLastPage
	}

	vm := view_models.NewPaginationVm()
	res = vm.Build(view_models.DetailPaginationVm{
		CurrentPage: page,
		LastPage:    lastPage,
		Total:       total,
		PerPage:     limit,
	})

	return res
}
