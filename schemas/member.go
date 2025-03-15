package schemas

import (
	"gorm.io/gorm"
)

type Member struct {
	gorm.Model
	Email           string `gorm:"type:varchar(255);unique" json:"email"`
	Token           string `gorm:"type:varchar(1024)" json:"token"`
	ProfileImageUrl string `gorm:"type:varchar(512)" json:"profile_image_url"`
	SnsType         string `gorm:"type:ENUM('KAKAO','NAVER','FACEBOOK','GOOGLE','APPLE'); DEFAULT:'GOOGLE'" json:"sns_type"`
	ActiveYn        bool   `gorm:"default:false" json:"active_yn"`
	UserLevel       int8   `gorm:"default:5" json:"user_level"`
	AllocatedDb     int8   `json:"allocted_db"`
}

func (Member) TableName() string {
	return "members"
}
