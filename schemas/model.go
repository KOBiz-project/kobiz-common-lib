package schemas

import (
	"gorm.io/gorm"
)

type Model struct {
	gorm.Model
	ModelName     string `gorm:"column:model_name;not null;type:varchar(256);comment:'모델 명'" json:"modelName"`
	Description   string `gorm:"column:description;type:text;comment:'모델 설명'" json:"description"`
	IdCompany     int    `gorm:"column:id_company;not null;comment:'제조사 고유 아이디'" json:"idCompany"`
	NameCompany   string `gorm:"column:name_company;not null;type:varchar(256);comment:'제조사 명'" json:"nameCompany"`
	Size64G       bool   `gorm:"column:size_64g;default:false;comment:'용량 64GB 여부'" json:"size64G"`
	Size64GCDSKT  string `gorm:"column:size_64g_cd_skt;type:varchar(256);comment:'용량 64GB SKT 모델 고유코드 (model_cd)'" json:"size64GCDSKT"`
	Size64GCDKT   string `gorm:"column:size_64g_cd_kt;type:varchar(256);comment:'용량 64GB KT 모델 고유코드 (prod_no)'" json:"size64GCDKT"`
	Size64GCDLGT  string `gorm:"column:size_64g_cd_lgt;type:varchar(256);comment:'용량 64GB LGT 모델 고유코드 (urc_trm_mdl_cd)'" json:"size64GCDLGT"`
	Size128G      bool   `gorm:"column:size_128g;default:false;comment:'용량 128GB 여부'" json:"size128G"`
	Size128GCDSKT string `gorm:"column:size_128g_cd_skt;type:varchar(256);comment:'용량 128GB SKT 모델 고유코드'" json:"size128GCDSKT"`
	Size128GCDKT  string `gorm:"column:size_128g_cd_kt;type:varchar(256);comment:'용량 128GB KT 모델 고유코드'" json:"size128GCDKT"`
	Size128GCDLGT string `gorm:"column:size_128g_cd_lgt;type:varchar(256);comment:'용량 128GB LGT 모델 고유코드'" json:"size128GCDLGT"`
	Size256G      bool   `gorm:"column:size_256g;default:false;comment:'용량 256GB 여부'" json:"size256G"`
	Size256GCDSKT string `gorm:"column:size_256g_cd_skt;type:varchar(256);comment:'용량 256GB SKT 모델 고유코드'" json:"size256GCDSKT"`
	Size256GCDKT  string `gorm:"column:size_256g_cd_kt;type:varchar(256);comment:'용량 256GB KT 모델 고유코드'" json:"size256GCDKT"`
	Size256GCDLGT string `gorm:"column:size_256g_cd_lgt;type:varchar(256);comment:'용량 256GB LGT 모델 고유코드'" json:"size256GCDLGT"`
	Size512G      bool   `gorm:"column:size_512g;default:false;comment:'용량 512GB 여부'" json:"size512G"`
	Size512GCDSKT string `gorm:"column:size_512g_cd_skt;type:varchar(256);comment:'용량 512GB SKT 모델 고유코드'" json:"size512GCDSKT"`
	Size512GCDKT  string `gorm:"column:size_512g_cd_kt;type:varchar(256);comment:'용량 512GB KT 모델 고유코드'" json:"size512GCDKT"`
	Size512GCDLGT string `gorm:"column:size_512g_cd_lgt;type:varchar(256);comment:'용량 512GB LGT 모델 고유코드'" json:"size512GCDLGT"`
	Size1T        bool   `gorm:"column:size_1t;default:false;comment:'용량 1TB 여부'" json:"size1T"`
	Size1TCDSKT   string `gorm:"column:size_1t_cd_skt;type:varchar(256);comment:'용량 1TB SKT 모델 고유코드'" json:"size1TCDSKT"`
	Size1TCDKT    string `gorm:"column:size_1t_cd_kt;type:varchar(256);comment:'용량 1TB KT 모델 고유코드'" json:"size1TCDKT"`
	Size1TCDLGT   string `gorm:"column:size_1t_cd_lgt;type:varchar(256);comment:'용량 1TB LGT 모델 고유코드'" json:"size1TCDLGT"`
	Color1        string `gorm:"column:color1;type:varchar(256);comment:'색상 1'" json:"color1"`
	Color1HexCode string `gorm:"column:color1_hex_code;type:varchar(256);comment:'색상 1 Hex Code'" json:"color1HexCode"`
	Color1Photo   string `gorm:"column:color1_photo;type:varchar(512);comment:'색상 1 폰 이미지'" json:"color1Photo"`
	Color2        string `gorm:"column:color2;type:varchar(256);comment:'색상 2'" json:"color2"`
	Color2HexCode string `gorm:"column:color2_hex_code;type:varchar(256);comment:'색상 2 Hex Code'" json:"color2HexCode"`
	Color2Photo   string `gorm:"column:color2_photo;type:varchar(512);comment:'색상 2 폰 이미지'" json:"color2Photo"`
	Color3        string `gorm:"column:color3;type:varchar(256);comment:'색상 3'" json:"color3"`
	Color3HexCode string `gorm:"column:color3_hex_code;type:varchar(256);comment:'색상 3 Hex Code'" json:"color3HexCode"`
	Color3Photo   string `gorm:"column:color3_photo;type:varchar(512);comment:'색상 3 폰 이미지'" json:"color3Photo"`
}

func (Model) TableName() string {
	return "models"
}
