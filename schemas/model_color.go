package schemas

import (
	"gorm.io/gorm"
)

type ModelColor struct {
	gorm.Model
	IdModel      int    `gorm:"column:id_model;not null;comment:'모델 고유 아이디'" json:"idModel"`
	Type         string `gorm:"column:type;not nulltype:varchar(20);comment:'통신사 구분'" json:"type"`
	ColorName    string `gorm:"column:name;type:varchar(256);comment:'색상'" json:"color"`
	ColorHexCode string `gorm:"column:hex_code;type:varchar(256);comment:'Hex Code'" json:"hexCode"`
	ColorPhoto   string `gorm:"column:photo;type:varchar(512);comment:'폰 이미지'" json:"photo"`
}

func (ModelColor) TableName() string {
	return "model_colors"
}
