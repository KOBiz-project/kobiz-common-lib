package schemas

import "gorm.io/gorm"

type MemberAdmin struct {
	gorm.Model
	UserId       string `gorm:"type:varchar(50);unique" json:"user_id"`
	Password     string `gorm:"type:varchar(1024)" json:"password"`
	ActiveYn     bool   `gorm:"default:false" json:"active_yn"`
	Name         string `gorm:"type:varchar(50)" json:"name"`
	NickName     string `gorm:"type:varchar(50);unique;column:nick_name" json:"nick_name"`
	UserLevel    int8   `gorm:"default:5" json:"user_level"`
	ProfileImage string `gorm:"type:varchar(1024)" json:"profile_image"`
	Creator      string `gorm:"type:varchar(50)" json:"creator"`
	Updator      string `gorm:"type:varchar(50)" json:"updator"`
}

func (MemberAdmin) TableName() string {
	return "member_admins"
}
