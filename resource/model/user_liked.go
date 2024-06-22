package model

import "time"

type UserLiked struct {
	Id         int    `json:"id"`
	Phone      string `json:"phone" gorm:"type:varchar(16); foreignKey:Phone;references:Phone"`
	PhoneLiked string `json:"phone_liked" gorm:"type:varchar(16); foreignKey:PhoneLiked;references:Phone"`
	CreatedAt  int64  `json:"-" gorm:"bigint"`
	UpdatedAt  int64  `json:"-" gorm:"bigint"`
}

func (u *UserLiked) SetTime() {
	u.CreatedAt = time.Now().Unix()
	u.UpdatedAt = time.Now().Unix()
}

func (u *UserLiked) UpdateTime() {
	u.UpdatedAt = time.Now().Unix()
}
