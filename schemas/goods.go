package schemas

import (
	"gorm.io/gorm"
)

// 자체 상품 스키마
type Goods struct {
	gorm.Model
	Telecom                     string `gorm:"column:telecom;not null;type:varchar(30);comment:'통신사'" json:"telecom"`
	IdModel                     int64  `gorm:"column:id_model;not null;comment:'모델 고유 ID'" json:"idModel"`
	CustomNameModel             string `gorm:"column:custom_name_model;type:varchar(256);comment:'커스텀 모델명'" json:"customNameModel"`
	CodeModel                   string `gorm:"column:code_model;type:varchar(256);comment:'모델 고유코드'" json:"codeModel"`
	Size                        string `gorm:"column:size;not null;type:varchar(10);comment:'기기용량 64GB/128GB/256GB/512GB/1TB'" json:"size"`
	NumChg                      bool   `gorm:"column:num_chg;default:false;comment:'번호 이동 여부'" json:"numChg"`
	NumChgIdPlan                int64  `gorm:"column:num_chg_id_plan;comment:'번호 이동 요금제 ID'" json:"numChgIdPlan"`
	NumChgCustomNamePlan        string `gorm:"column:num_chg_custom_name_plan;type:varchar(256);comment:'번호 이동 커스텀 요금제명';" json:"numChgCustomNamePlan"`
	NumChgSwitchSubsidy         int64  `gorm:"column:num_chg_switch_subsidy;type:bigint;comment:'번호 이동 전환지원금'" json:"numChgSwitchSubsidy"`
	NumChgAddSubsidy            int64  `gorm:"column:num_chg_add_subsidy;type:bigint;comment:'번호 이동 추가지원금'" json:"numChgAddSubsidy"`
	DeviceChg                   bool   `gorm:"column:device_chg;default:false;comment:'기기 변경 여부'" json:"deviceChg"`
	DeviceChgCustomNamePlan     string `gorm:"column:device_chg_custom_name_plan;type:varchar(256);comment:'기기 변경 커스텀 요금제명';" json:"deviceChgCustomNamePlan"`
	DeviceChgIdPlan             int64  `gorm:"column:device_chg_id_plan;comment:'기기 변경 요금제 ID'" json:"deviceChgIdPlan"`
	DeviceChgAddSubsidy         int64  `gorm:"column:device_chg_add_subsidy;type:bigint;comment:'기기 변경 추가지원금'" json:"deviceChgAddSubsidy"`
	NewActivation               bool   `gorm:"column:new_activation;default:false;comment:'신규 개통 여부'" json:"newActivation"`
	NewActivationIdPlan         int64  `gorm:"column:new_activation_id_plan;comment:'신규 개통 요금제 ID'" json:"newActivationIdPlan"`
	NewActivationCustomNamePlan string `gorm:"column:new_activation_custom_name_plan;type:varchar(256);comment:'신규 개통 커스텀 요금제명';" json:"newActivationCustomNamePlan"`
	NewActivationAddSubsidy     int64  `gorm:"column:new_activation_add_subsidy;type:bigint;comment:'신규 개통 추가지원금'" json:"newActivationAddSubsidy"`
	SubsidyYn                   bool   `gorm:"column:subsidy_yn;default:false;comment:'공시지원 여부'" json:"subsidyYn"`
	Description                 string `gorm:"column:description;type:text;comment:'상품 설명'" json:"description"`
	ActiveYn                    bool   `gorm:"column:active_yn;default:false" json:"activeYn"`
}

func (Goods) TableName() string {
	return "goods"
}
