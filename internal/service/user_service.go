package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/farhaniupr/dating-api/internal/helper"
	"github.com/farhaniupr/dating-api/internal/repository"
	"github.com/farhaniupr/dating-api/package/eksternal"
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
	FindDate(ctx context.Context, jwtUser map[string]interface{}) (result model.User, err error)
}

type UserService struct {
	env            library.Env
	repositoryUser repository.UserRepository
	jwtService     JWTAuthMethodService
	commonHelper   helper.CommonHelper
	encryptHelper  helper.EcncryptHelper
	redisCrud      eksternal.RedisEksternal
}

func ModuleUserService(
	env library.Env,
	repositoryUser repository.UserRepository,
	commonHelper helper.CommonHelper,
	jwtService JWTAuthMethodService,
	encryptHelper helper.EcncryptHelper,
	redisCrud eksternal.RedisEksternal,

) UserMethodService {
	return UserService{
		env:            env,
		repositoryUser: repositoryUser,
		commonHelper:   commonHelper,
		jwtService:     jwtService,
		encryptHelper:  encryptHelper,
		redisCrud:      redisCrud,
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
		token, err := u.jwtService.CreateToken(resultUser)
		if err != nil {
			return false, model.User{}, err
		}

		resultUser.Token = token
		resultUser.Password = ""

		return true, resultUser, nil
	}

	return false, model.User{}, errors.New("password is wrong")
}

func (u UserService) DetailUser(ctx context.Context, phone string) (user model.User, err error) {

	return u.repositoryUser.WithContext(ctx).DetailUser(phone)
}

func (u UserService) FindDate(ctx context.Context, jwtUser map[string]interface{}) (user model.User, err error) {
	var dataExpect []string

	resultKey := u.redisCrud.GetList(fmt.Sprintf("finddate/%s/*", jwtUser["phone"]))

	if len(resultKey) >= 10 {
		return model.User{}, errors.New("out of limit find date")
	}

	for _, value := range resultKey {
		dataExpect = append(dataExpect, u.redisCrud.Get(value))
	}

	user, err = u.repositoryUser.WithContext(ctx).FindUser(jwtUser["phone"], jwtUser["gender"], dataExpect)
	if err != nil {
		return model.User{}, err
	}

	user.Password = ""

	u.redisCrud.Store1Day(fmt.Sprintf("finddate/%s/%s", jwtUser["phone"], user.Phone), user.Phone)

	return user, err
}

func (u UserService) StoreUser(tx *gorm.DB, dataReq model.User) (user model.User, err error) {

	dataReq.Subscription = "free"
	dataReq.Verify = "no"

	dataReq.Password, err = u.encryptHelper.EncryptPassword(dataReq.Password)
	if err != nil {
		return model.User{}, err
	}

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
