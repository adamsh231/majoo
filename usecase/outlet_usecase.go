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

type OutletUseCase struct {
	*Contract
}

func NewOutletUseCase(ucContract *Contract) interfaces.IOutletUseCase {
	return &OutletUseCase{Contract: ucContract}
}

func (uc OutletUseCase) Browse(search, orderBy, sort string, page, limit int) (res []*view_models.OutletListVM, pagination view_models.PaginationVm, err error) {
	repository := repositories.NewOutletRepository(uc.PostgresDB)
	offset, limit, page, orderBy, sort := uc.setPaginationParameter(page, limit, orderBy, sort)

	outlets, err := repository.Browse(search, orderBy, sort, limit, offset)
	if err != nil {
		helper.LogOnly(err.Error(), "uc-browse")
		return res, pagination, err
	}
	for _, outlet := range outlets {
		vm := view_models.NewOutletListVM()
		vm.Build(&outlet)
		res = append(res, vm)
	}

	//set pagination
	totalCount, err := repository.Count(search)
	if err != nil {
		helper.LogOnly(err.Error(), "uc-browse-count")
		return res, pagination, err
	}
	pagination = uc.setPaginationResponse(page, limit, totalCount)

	return res, pagination, nil
}

func (uc OutletUseCase) Read(id string) (res view_models.OutletDetailVM, err error) {
	panic("implement me")
}

func (uc OutletUseCase) Add(req *requests.OutletAddRequest) (res string, err error) {
	now := time.Now().UTC()
	model := models.Outlet{
		Merchant: models.Merchant{ID: req.MerchantID},
		Name:      req.Name,
		Phone:     req.Phone,
		Address:   req.Address,
		CreatedAt: now,
		UpdatedAt: now,
	}

	repo := repositories.NewOutletRepository(uc.PostgresDB)
	res, err = repo.Add(model, uc.PostgresTX)
	if err != nil {
		helper.LogOnly(err.Error(), "uc-add")
		return res, err
	}

	return res, err
}

func (uc OutletUseCase) Edit(req *requests.OutletEditRequest, id string) (res string, err error) {
	panic("implement me")
}

func (uc OutletUseCase) Delete(id string) (res string, err error) {
	panic("implement me")
}
