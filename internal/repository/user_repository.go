package repository

import (
	"context"
	"log"

	"github.com/farhaniupr/dating-api/package/library"
	"github.com/farhaniupr/dating-api/resource/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	db library.Database
}

func ModuleUserRepository(db library.Database) UserRepository {
	return UserRepository{
		db: db,
	}
}

func (u UserRepository) WithTransaction(txHandle *gorm.DB) UserRepository {
	if txHandle == nil {
		log.Println("not found transaction context")
		return u
	}
	u.db.MysqlDB = txHandle
	return u

}

func (u UserRepository) WithContext(ctx context.Context) UserRepository {
	u.db.MysqlDB = u.db.MysqlDB.WithContext(ctx)
	return u
}

func (u UserRepository) DetailUser(phone string) (result model.User, err error) {
	return result, u.db.MysqlDB.Debug().Table("user").Where("phone = ?", phone).
		Select("*").Scan(&result).Error
}

func (u UserRepository) FindUser(phone, gender interface{}, except []string) (result model.User, err error) {

	query := u.db.MysqlDB.Table("user u").
		Joins("left join user_liked us on u.phone = us.phone_liked and us.phone = ?", phone).
		Limit(1).
		Order("rand()").
		Where("u.phone != ?", phone).
		Where("u.gender != ?", gender).
		Select("u.*, case when us.id > 0 then 'liked' else  'ready liked' end as 'status_like'")

	if len(except) > 0 {
		query = query.Where("u.phone not in ?", except)
	}

	return result, query.Scan(&result).Error
}

func (u UserRepository) CreateUser(dataReq model.User) (result model.User, err error) {
	return result, u.db.MysqlDB.
		Omit("status_like").
		Create(&dataReq).Scan(&result).Error
}

func (u UserRepository) UpdateUser(dataReq model.User, phone string) (result model.User, rowsAffected int, err error) {
	query := u.db.MysqlDB.Table("user").
		Omit("status_like").
		Where("phone = ?", phone).Updates(&dataReq).Scan(&result)

	return result, int(query.RowsAffected), err
}

func (u UserRepository) DeleteUser(dataReq model.User, phone string) (result model.User, rowsAffected int, err error) {
	query := u.db.MysqlDB.Table("user").Where("phone = ?", phone).Delete(&dataReq)
	return result, int(query.RowsAffected), query.Error
}

func (u UserRepository) ListUser(limit, page int) (result []model.User, total int64, err error) {
	u.db.MysqlDB.Table("user").
		Count(&total)

	return result, total, u.db.MysqlDB.Debug().Table("user").
		Offset((page - 1) * limit).Limit(limit).Select("*").
		Scan(&result).Error
}
