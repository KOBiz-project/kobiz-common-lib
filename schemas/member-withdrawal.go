package schemas

import (
	"gorm.io/gorm"
)

type MemberWithdrawal struct {
	gorm.Model
	IdGroup      int    `gorm:"column:id_group;not null;comment:'회원 그룹 아이디'" json:"idGroup"`
	Email        string `gorm:"column:email;type:varchar(255);comment:'가입 이메일';uniqueIndex:idx_user" json:"email"`
	SnsType      string `gorm:"column:sns_type;type:ENUM('KAKAO','GOOGLE'); default:'GOOGLE';comment:'SNS 로그인 타입';uniqueIndex:idx_user" json:"snsType"`
	UserType     int    `gorm:"column:user_type;default:0;comment:'회원 유형 0: 미선택 1:기업회원 2: 일반회원'" json:"userType"`
	UserDocument string `gorm:"column:user_document;type:varchar(512);comment:'가입 시 제출서류'" json:"userDocument"`
	Memo1        string `gorm:"column:memo_1;type:varchar(255);comment:'짧은 메모 (회사명)'" json:"memo1"`
	Memo2        string `gorm:"column:memo_2;type:text;comment:'긴 메모 (관리자 메모)'" json:"memo2"`
}

func (MemberWithdrawal) TableName() string {
	return "members_withdrawal"
}
