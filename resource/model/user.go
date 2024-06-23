package model

import "time"

type User struct {
	Phone          string      `json:"phone" gorm:"type:varchar(16);primary_key" validate:"required"`
	Email          string      `json:"email" gorm:"type:varchar(255);uniqueIndex" validate:"required"`
	Name           string      `json:"name" gorm:"type:varchar(255)" validate:"required"`
	Password       string      `json:"password,omitempty" gorm:"type:varchar(255)" validate:"required"`
	UrlPhoto       string      `json:"url_photo" gorm:"type:varchar(255)" validate:"required"`
	DateBirth      string      `json:"date_birth" gorm:"type:datetime" validate:"required,date"`
	Gender         string      `json:"gender" gorm:"type:enum('male','female','')" validate:"required"`
	AboutMe        string      `json:"about_me" gorm:"type:longtext"`
	InstragramUrl  string      `json:"instragram_url" gorm:"type:varchar(255)"`
	City           string      `json:"city" gorm:"type:varchar(255)" validate:"required"`
	Country        string      `json:"country" gorm:"type:varchar(255)" validate:"required"`
	Subscription   string      `json:"subscription" gorm:"type:enum('free','premium') default 'free'"`
	Verify         string      `json:"verify" gorm:"type:enum('yes','no') default 'no'"`
	CreatedAt      int64       `json:"-" gorm:"type:bigint"`
	UpdatedAt      int64       `json:"-" gorm:"type:bigint"`
	UserLiked      UserLiked   `json:"-" gorm:"foreignKey:phone;references:phone"`
	UserSetting    UserSetting `json:"-" gorm:"foreignKey:phone;references:phone"`
	UserLikedPhone UserLiked   `json:"-" gorm:"foreignKey:phone_liked;references:phone"`
	StatusLike     string      `json:"status_like"`
	Token          string      `json:"token,omitempty" gorm:"-"`
}

func (u *User) SetTime() {
	u.CreatedAt = time.Now().Unix()
	u.UpdatedAt = time.Now().Unix()
}

func (u *User) UpdateTime() {
	u.UpdatedAt = time.Now().Unix()
}
