package schemas

import (
	"time"

	"gorm.io/gorm"
)

type MemberDetail struct {
	gorm.Model
	IDMember         int64     `gorm:"unique" json:"id_member"`
	Email            string    `gorm:"type:varchar(1024)" json:"email"`
	Name             string    `gorm:"type:varchar(50)" json:"name"`
	NickName         string    `gorm:"type:varchar(50);unique;column:nick_name" json:"nick_name"`
	ProfilePhoto     string    `gorm:"type:varchar(1024)" json:"profile_photo"`
	Tel              string    `gorm:"type:varchar(50)" json:"tel"`
	MobileCompany    int8      `gorm:"default:0" json:"mobile_company"`
	Mobile           string    `gorm:"type:varchar(50)" json:"mobile"`
	Address          string    `gorm:"type:varchar(1024)" json:"address"`
	AddressDetail    string    `gorm:"type:varchar(1024)" json:"address_detail"`
	Zipcode          string    `gorm:"type:varchar(50)" json:"zipcode"`
	AuthenticationCi string    `gorm:"type:varchar(255)" json:"authentication_ci"`
	AuthenticationAt time.Time `json:"authentication_at"`
	AltNewEvent      bool      `gorm:"default:false" json:"alt_new_event"`
	AltSuccessfulBid bool      `gorm:"default:false" json:"alt_successful_bid"`
	AltNewContent    bool      `gorm:"default:false" json:"alt_new_content"`
	AltNightPush     bool      `gorm:"default:false" json:"alt_night_push"`
	JoinScore        int64     `gorm:"default:0" json:"join_score"`
	Like             int64     `gorm:"default:0" json:"like"`
	Hate             int64     `gorm:"default:0" json:"hate"`
}

func (MemberDetail) TableName() string {
	return "member_details"
}
