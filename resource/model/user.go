package model

import "time"

type User struct {
	Phone          string      `json:"phone" gorm:"type:varchar(16);primary_key"`
	Email          string      `json:"email" gorm:"type:varchar(255)"`
	Name           string      `json:"name" gorm:"type:varchar(255)"`
	Password       string      `json:"password" gorm:"type:varchar(255)"`
	UrlPhoto       string      `json:"url_photo" gorm:"type:varchar(255)"`
	DateBirth      string      `json:"date_birth" gorm:"type:datetime"`
	Gender         string      `json:"gender" gorm:"type:enum('male','female','')"`
	AboutMe        string      `json:"about_me" gorm:"type:longtext"`
	InstragramUrl  string      `json:"instragram_url" gorm:"type:varchar(255)"`
	City           string      `json:"city" gorm:"type:varchar(255)"`
	Country        string      `json:"country" gorm:"type:varchar(255)"`
	Subscription   string      `json:"subscription" gorm:"type:enum('free','premium') default 'free'"`
	Verify         string      `json:"verify" gorm:"type:enum('yes','no') default 'no'"`
	CreatedAt      int64       `json:"-" gorm:"type:bigint"`
	UpdatedAt      int64       `json:"-" gorm:"type:bigint"`
	UserSetting    UserSetting `json:"user_setting" gorm:"foreignKey:phone;references:phone"`
	UserLiked      UserLiked   `json:"-" gorm:"foreignKey:phone;references:phone"`
	UserLikedPhone UserLiked   `json:"-" gorm:"foreignKey:phone_liked;references:phone"`
	Token          string      `json:"token" gorm:"-"`
}

func (u *User) SetTime() {
	u.CreatedAt = time.Now().Unix()
	u.UpdatedAt = time.Now().Unix()
}

func (u *User) UpdateTime() {
	u.UpdatedAt = time.Now().Unix()
}
