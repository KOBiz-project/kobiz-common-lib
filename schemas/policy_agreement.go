package schemas

import (
	"gorm.io/gorm"
)

type PolicyAgreement struct {
	gorm.Model
	IdMember int `gorm:"column:id_member;not null;comment:'회원 고유 아이디'" json:"idMember"`
	IdPolicy int `gorm:"column:id_policy;not null;comment:'약관 고유 아이디 - 1 : 서비스 이용약관 2: 개인정보 처리방침 3: 개인정보 수집 및 이용 동의 4: 광고성 정보 수신 동의 (선택)'" json:"idPolicy"`
}

func (PolicyAgreement) TableName() string {
	return "policy_agreement"
}
