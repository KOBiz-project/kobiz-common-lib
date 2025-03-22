package schemas

import (
	"gorm.io/gorm"
)

type Member struct {
	gorm.Model
	Email           string `gorm:"type:varchar(255);unique" json:"email"`
	Token           string `gorm:"type:varchar(1024)" json:"token"`
	ProfileImageUrl string `gorm:"type:varchar(512)" json:"profileImageUrl"`
	SnsType         string `gorm:"type:ENUM('KAKAO','NAVER','FACEBOOK','GOOGLE','APPLE'); DEFAULT:'GOOGLE'" json:"snsType"`
	ActiveYn        bool   `gorm:"default:false" json:"activeYn"`
	UserLevel       int8   `gorm:"default:5" json:"userLevel"`
}

func (Member) TableName() string {
	return "members"
}
