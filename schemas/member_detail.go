package schemas

import (
	"time"

	"gorm.io/gorm"
)

type MemberDetail struct {
	gorm.Model
	IDMember         int64     `gorm:"unique" json:"idMember"`
	Email            string    `gorm:"type:varchar(1024)" json:"email"`
	Name             string    `gorm:"type:varchar(50)" json:"name"`
	NickName         string    `gorm:"type:varchar(50);unique;column:nick_name" json:"nickName"`
	ProfilePhoto     string    `gorm:"type:varchar(1024)" json:"profilePhoto"`
	Tel              string    `gorm:"type:varchar(50)" json:"tel"`
	MobileCompany    int8      `gorm:"default:0" json:"mobileCompany"`
	Mobile           string    `gorm:"type:varchar(50)" json:"mobile"`
	Address          string    `gorm:"type:varchar(1024)" json:"address"`
	AddressDetail    string    `gorm:"type:varchar(1024)" json:"addressDetail"`
	Zipcode          string    `gorm:"type:varchar(50)" json:"zipcode"`
	AuthenticationCi string    `gorm:"type:varchar(255)" json:"authenticationCi"`
	AuthenticationAt time.Time `json:"authenticationAt"`
}

func (MemberDetail) TableName() string {
	return "member_details"
}
