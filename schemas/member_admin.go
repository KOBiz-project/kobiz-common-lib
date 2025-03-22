package schemas

import "gorm.io/gorm"

type MemberAdmin struct {
	gorm.Model
	UserId       string `gorm:"type:varchar(50);unique" json:"userId"`
	Password     string `gorm:"type:varchar(1024)" json:"password"`
	ActiveYn     bool   `gorm:"default:false" json:"activeYn"`
	Name         string `gorm:"type:varchar(50)" json:"name"`
	NickName     string `gorm:"type:varchar(50);unique;column:nick_name" json:"nickName"`
	UserLevel    int8   `gorm:"default:5" json:"userLevel"`
	ProfileImage string `gorm:"type:varchar(1024)" json:"profileImage"`
}

func (MemberAdmin) TableName() string {
	return "member_admins"
}
