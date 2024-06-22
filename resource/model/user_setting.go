package model

type UserSetting struct {
	Id             int    `json:"id"`
	Phone          string `json:"phone" gorm:"type:varchar(16)"`
	SexualInterest string `json:"sexual_interest" gorm:"type:enum('male', 'female','all') default 'all'"`
}
