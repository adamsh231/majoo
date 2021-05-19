package usecase

import (
	"github.com/adamsh231/majoo/domain/interfaces"
	"github.com/adamsh231/majoo/domain/models"
	"github.com/adamsh231/majoo/domain/requests"
	"github.com/adamsh231/majoo/domain/view_models"
	"github.com/adamsh231/majoo/packages/helper"
	"github.com/adamsh231/majoo/repositories"
	"time"
)

type ProductImageUseCase struct {
	*Contract
}

func NewProductImageUseCase(ucContract *Contract) interfaces.IProductImageUseCase {
	return &ProductImageUseCase{Contract: ucContract}
}

func (uc ProductImageUseCase) Browse(search, orderBy, sort string, page, limit int) (res []*view_models.ProductImageListVM, pagination view_models.PaginationVm, err error) {
	panic("implement me")
}

func (uc ProductImageUseCase) Read(id string) (res view_models.ProductImageDetailVM, err error) {
	panic("implement me")
}

func (uc ProductImageUseCase) Add(req *requests.ProductImageAddRequest, fileName string) (res string, err error) {
	now := time.Now().UTC()
	model := models.ProductImage{
		ProductID: req.ProductID,
		Path:      fileName,
		Alt:       req.Alt,
		CreatedAt: now,
		UpdatedAt: now,
	}

	repo := repositories.NewProductImageRepository(uc.PostgresDB)
	res, err = repo.Add(model, uc.PostgresTX)
	if err != nil {
		helper.LogOnly(err.Error(), "uc-add")
		return res, err
	}

	return res, err
}

func (uc ProductImageUseCase) Edit() (res string, err error) {
	panic("implement me")
}

func (uc ProductImageUseCase) Delete(id string) (res string, err error) {
	panic("implement me")
}