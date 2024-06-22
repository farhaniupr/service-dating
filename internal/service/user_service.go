package service

import (
	"context"
	"errors"

	"github.com/farhaniupr/dating-api/internal/helper"
	"github.com/farhaniupr/dating-api/internal/repository"
	"github.com/farhaniupr/dating-api/package/library"
	"github.com/farhaniupr/dating-api/resource/model"
	"gorm.io/gorm"
)

type UserMethodService interface {
	Login(context context.Context, dataReq model.User) (result bool, resultUser model.User, err error)
	DetailUser(ctx context.Context, id string) (result model.User, err error)
	ListUser(ctx context.Context, page, limit int) (result []model.User, total int64, err error)
	StoreUser(tx *gorm.DB, dataReq model.User) (result model.User, err error)
	UpdateUser(tx *gorm.DB, dataReq model.User, id string) (result model.User, err error)
	DeleteUser(ctx context.Context, dataReq model.User, id string) (result []interface{}, err error)
}

type UserService struct {
	env            library.Env
	repositoryUser repository.UserRepository
	jwtService     JWTAuthMethodService
	commonHelper   helper.CommonHelper
	encryptHelper  helper.EcncryptHelper
}

func ModuleUserService(
	env library.Env,
	repositoryUser repository.UserRepository,
	commonHelper helper.CommonHelper,
	jwtService JWTAuthMethodService,
	encryptHelper helper.EcncryptHelper,

) UserMethodService {
	return UserService{
		env:            env,
		repositoryUser: repositoryUser,
		commonHelper:   commonHelper,
		jwtService:     jwtService,
		encryptHelper:  encryptHelper,
	}
}

func (u UserService) Login(context context.Context, dataReq model.User) (result bool, resultUser model.User, err error) {

	resultUser, err = u.repositoryUser.WithContext(context).DetailUser(dataReq.Phone)
	if err != nil {
		return false, model.User{}, err
	}

	if resultUser.Phone == "" {
		return false, model.User{}, errors.New("account not found")
	}

	if u.encryptHelper.CheckPassword(resultUser.Password, dataReq.Password) {

		token, err := u.jwtService.CreateToken(dataReq)
		if err != nil {
			return false, model.User{}, err
		}

		resultUser.Token = token

		return true, resultUser, nil
	}

	return false, model.User{}, errors.New("password is wrong")
}

func (u UserService) DetailUser(ctx context.Context, id string) (user model.User, err error) {

	return u.repositoryUser.WithContext(ctx).DetailUser(id)
}

func (u UserService) StoreUser(tx *gorm.DB, dataReq model.User) (user model.User, err error) {
	return u.repositoryUser.WithTransaction(tx).CreateUser(dataReq)
}

func (u UserService) UpdateUser(tx *gorm.DB, dataReq model.User, id string) (user model.User, err error) {

	user, rowAffected, err := u.repositoryUser.WithTransaction(tx).UpdateUser(dataReq, id)

	if rowAffected == 0 {
		err = errors.New("data has not updated")
	}

	return user, err
}

func (u UserService) DeleteUser(ctx context.Context, dataReq model.User, id string) (user []interface{}, err error) {

	_, rowAffected, err := u.repositoryUser.WithContext(ctx).DeleteUser(dataReq, id)

	if rowAffected == 0 {
		err = errors.New("data has not deleted")
	}

	return user, err
}

func (u UserService) ListUser(ctx context.Context, page, limit int) (result []model.User, total int64, err error) {
	return u.repositoryUser.WithContext(ctx).ListUser(limit, page)
}
