package dbmanager

import (
	"time"

	"gorm.io/gorm"
)

type MemberAdminLoginLog struct {
	gorm.Model
	SeqMemberAdminLoginLog int64  `gorm:"primaryKey;autoIncrement:true" json:"seqMemberAdminLoginLog"`
	SeqMemberAdmin         int64  `gorm:"index" json:"seqMemberAdmin"`
	Token                  string `gorm:"type:varchar(1024)" json:"token"`
	LoginAt                time.Time
}

func (MemberAdminLoginLog) TableName() string {
	return "member_admin_login_logs"
}
