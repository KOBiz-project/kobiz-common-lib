package schemas

import (
	"time"

	"gorm.io/gorm"
)

type MemberAdminLoginLog struct {
	gorm.Model
	SeqMemberAdminLoginLog int64  `gorm:"primaryKey;autoIncrement:true" json:"seq_member_admin_login_log"`
	SeqMemberAdmin         int64  `gorm:"index" json:"seq_member_admin"`
	Token                  string `gorm:"type:varchar(1024)" json:"token"`
	LoginAt                time.Time
}

func (MemberAdminLoginLog) TableName() string {
	return "member_admin_login_logs"
}
