package schemas

import (
	"gorm.io/gorm"
)

type Member struct {
	gorm.Model
	Email    string `gorm:"type:varchar(255);comment:'가입 이메일';uniqueIndex:idx_user" json:"email"`
	SnsType  string `gorm:"type:ENUM('KAKAO','GOOGLE'); default:'GOOGLE';comment:'SNS 로그인 타입';uniqueIndex:idx_user" json:"snsType"`
	UserType int    `gorm:"column:user_type;default:0;comment:'회원 유형 0: 미선택 1:기업회원 2: 일반회원'" json:"userType"`
	ActiveYn bool   `gorm:"default:false;comment:'관리자 승인 여부'" json:"activeYn"`
}

func (Member) TableName() string {
	return "members"
}
