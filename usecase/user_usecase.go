package usecase

import (
	"database/sql"
	"errors"
	"github.com/adamsh231/majoo/domain/interfaces"
	"github.com/adamsh231/majoo/domain/models"
	"github.com/adamsh231/majoo/domain/requests"
	"github.com/adamsh231/majoo/domain/view_models"
	"github.com/adamsh231/majoo/packages/helper"
	"github.com/adamsh231/majoo/packages/messages"
	"github.com/adamsh231/majoo/repositories"
	"time"
)

type UserUseCase struct {
	*Contract
}

func NewUserUseCase(ucContract *Contract) interfaces.IUserUseCase {
	return &UserUseCase{Contract: ucContract}
}

func (uc UserUseCase) Login(req *requests.UserLoginRequest) (res view_models.LoginVm, err error) {
	model := models.User{Email: req.Email}
	repo := repositories.NewUserRepository(uc.PostgresDB)
	model, err = repo.Read(model)
	if err != nil {
		helper.LogOnly(err.Error(), "uc-login-read")
	}

	// check email
	if model.Email != req.Email {
		helper.LogOnly(messages.CredentialDoNotMatch, "uc-login-email")
		return res, errors.New(messages.CredentialDoNotMatch)
	}

	// check password
	if !helper.CheckHashString(req.Password, model.Password) {
		helper.LogOnly(messages.CredentialDoNotMatch, "uc-login-password")
		return res, errors.New(messages.CredentialDoNotMatch)
	}

	// generate jwt payload and encrypted with jwe
	payload := map[string]interface{}{
		"id":     model.ID,
		"role": model.Role.String,
	}
	jwePayload, err := uc.JweCredential.GenerateJwePayload(payload)
	if err != nil {
		helper.LogOnly(err.Error(), "uc-login-generate-jwt-payload")
		return res, err
	}

	// generate jwt token
	res, err = uc.generateJWT(req.Email, jwePayload)
	if err != nil {
		helper.LogOnly(messages.CredentialDoNotMatch, "uc-login-generate-jwt-token")
		return res, err
	}

	return res, err
}

func (uc UserUseCase) generateJWT(issuer, payload string) (res view_models.LoginVm, err error) {
	res.Token, res.TokenExpiredAt, err = uc.JwtCredential.GetToken(issuer, payload)
	if err != nil {
		return res, err
	}

	res.RefreshToken, res.RefreshTokenExpiredAt, err = uc.JwtCredential.GetRefreshToken(issuer, payload)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (uc UserUseCase) Add(req *requests.UserAddRequest) (res string, err error) {
	now := time.Now().UTC()
	passwordHash, err := helper.HashAndSalt(req.Password)
	if err != nil {
		helper.LogOnly(err.Error(), "uc-add-hash")
		return res, err
	}

	model := models.User{
		Name:      req.Name,
		Email:     req.Email,
		Password:  passwordHash,
		CreatedAt: now,
		UpdatedAt: now,
	}

	repo := repositories.NewUserRepository(uc.PostgresDB)
	res, err = repo.Add(model, uc.PostgresTX)
	if err != nil {
		helper.LogOnly(err.Error(), "uc-add")
		return res, err
	}

	return res, err
}

func (uc UserUseCase) Edit(req *requests.UserEditRequest, id string) (res string, err error) {
	now := time.Now().UTC()
	passwordHash, err := helper.HashAndSalt(req.Password)
	if err != nil {
		helper.LogOnly(err.Error(), "uc-edit-hash")
		return res, err
	}

	model := models.User{
		ID:        id,
		Name:      req.Name,
		Email:     req.Email,
		Password:  passwordHash,
		UpdatedAt: now,
	}

	repo := repositories.NewUserRepository(uc.PostgresDB)
	res, err = repo.Edit(model, uc.PostgresTX)
	if err != nil {
		helper.LogOnly(err.Error(), "uc-edit")
		return res, err
	}

	return res, err
}

func (uc UserUseCase) Delete(id string) (res string, err error) {
	now := time.Now().UTC()
	model := models.User{
		ID:        id,
		DeletedAt: sql.NullTime{Time: now, Valid: true},
	}

	repo := repositories.NewUserRepository(uc.PostgresDB)
	res, err = repo.Delete(model, uc.PostgresTX)
	if err != nil {
		helper.LogOnly(err.Error(), "uc-delete")
		return res, err
	}

	return res, err
}

