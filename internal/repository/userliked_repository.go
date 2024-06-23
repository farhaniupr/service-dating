package repository

import (
	"context"
	"log"

	"github.com/farhaniupr/dating-api/package/library"
	"github.com/farhaniupr/dating-api/resource/model"
	"gorm.io/gorm"
)

type UserLikedRepository struct {
	db library.Database
}

func ModuleUserLiked(db library.Database) UserLikedRepository {
	return UserLikedRepository{
		db: db,
	}
}

func (u UserLikedRepository) WithTransaction(txHandle *gorm.DB) UserLikedRepository {
	if txHandle == nil {
		log.Println("not found transaction context")
		return u
	}
	u.db.MysqlDB = txHandle
	return u

}

func (u UserLikedRepository) WithContext(ctx context.Context) UserLikedRepository {
	u.db.MysqlDB = u.db.MysqlDB.WithContext(ctx)
	return u
}

func (u UserLikedRepository) Create(dataReq model.UserLiked) (result model.UserLiked, err error) {
	return result, u.db.MysqlDB.Table("user_liked").Create(&dataReq).Scan(&result).Error
}

func (u UserLikedRepository) Detail(dataReq model.UserLiked) (result model.UserLiked, err error) {
	return result, u.db.MysqlDB.Table("user_liked").Where("phone = ?", dataReq.Phone).Where("phone_liked = ?", dataReq.PhoneLiked).Scan(&result).Scan(&result).Error
}
