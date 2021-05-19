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

type MerchantUseCase struct {
	*Contract
}

func NewMerchantUseCase(ucContract *Contract) interfaces.IMerchantUseCase {
	return &MerchantUseCase{Contract: ucContract}
}

func (uc MerchantUseCase) Browse(search, orderBy, sort string, page, limit int) (res []*view_models.MerchantListVM, pagination view_models.PaginationVm, err error) {
	repository := repositories.NewMerchantRepository(uc.PostgresDB)
	offset, limit, page, orderBy, sort := uc.setPaginationParameter(page, limit, orderBy, sort)

	merchants, err := repository.Browse(search, orderBy, sort, limit, offset)
	if err != nil {
		helper.LogOnly(err.Error(), "uc-browse")
		return res, pagination, err
	}
	for _, merchant := range merchants {
		vm := view_models.NewMerchantListVM()
		vm.Build(&merchant)
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

func (uc MerchantUseCase) Read(id string) (res view_models.MerchantDetailVM, err error) {
	panic("implement me")
}

func (uc MerchantUseCase) Add(req *requests.MerchantAddRequest) (res string, err error) {
	now := time.Now().UTC()
	model := models.Merchant{
		Name:      req.Name,
		Phone:     req.Phone,
		Address:   req.Address,
		CreatedAt: now,
		UpdatedAt: now,
	}

	repo := repositories.NewMerchantRepository(uc.PostgresDB)
	res, err = repo.Add(model, uc.PostgresTX)
	if err != nil {
		helper.LogOnly(err.Error(), "uc-add")
		return res, err
	}

	return res, err
}

func (uc MerchantUseCase) Edit(req *requests.MerchantEditRequest, id string) (res string, err error) {
	panic("implement me")
}

func (uc MerchantUseCase) Delete(id string) (res string, err error) {
	panic("implement me")
}
