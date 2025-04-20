package schemas

import (
	"gorm.io/gorm"
)

// 자체 상품 스키마
type Goods struct {
	gorm.Model
	IdModel                 int    `gorm:"column:id_model;not null;comment:'모델 고유 아이디'" json:"idModel"`
	Size                    string `gorm:"column:size;not null;type:varchar(10);comment:'기기용량 64GB/128GB/256GB/512GB/1TB'" json:"size"`
	NumChg                  bool   `gorm:"column:num_chg;default:false;comment:'번호 이동 여부'" json:"numChg"`
	NumChgCodePlan          string `gorm:"column:num_chg_code_plan;type:varchar(256);comment:'번호 이동 요금제 코드'" json:"numChgCodePlan"`
	NumChgSwitchSubsidy     int64  `gorm:"column:num_chg_switch_subsidy;type:bigint;comment:'번호 이동 전환지원금'" json:"numChgSwitchSubsidy"`
	NumChgAddSubsidy        int64  `gorm:"column:num_chg_add_subsidy;type:bigint;comment:'번호 이동 추가지원금'" json:"numChgAddSubsidy"`
	DeviceChg               bool   `gorm:"column:device_chg;default:false;comment:'기기 변경 여부'" json:"deviceChg"`
	DeviceChgCodePlan       string `gorm:"column:device_chg_code_plan;type:varchar(256);comment:'기기 변경 요금제 코드'" json:"deviceChgCodePlan"`
	DeviceChgAddSubsidy     int64  `gorm:"column:device_chg_add_subsidy;type:bigint;comment:'기기 변경 추가지원금'" json:"deviceChgAddSubsidy"`
	NewActivation           bool   `gorm:"column:new_activation;default:false;comment:'신규 개통 여부'" json:"newActivation"`
	NewActivationCodePlan   string `gorm:"column:new_activation_code_plan;type:varchar(256);comment:'신규 개통 요금제 코드'" json:"newActivationCodePlan"`
	NewActivationAddSubsidy int64  `gorm:"column:new_activation_add_subsidy;type:bigint;comment:'신규 개통 추가지원금'" json:"newActivationAddSubsidy"`
	SubsidyYn               bool   `gorm:"column:subsidy_yn;default:false;comment:'공시지원 여부'" json:"subsidyYn"`
	ActiveYn                bool   `gorm:"column:active_yn;default:false" json:"activeYn"`
}

func (Goods) TableName() string {
	return "goods"
}
