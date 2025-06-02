package schemas

import (
	"time"

	"gorm.io/gorm"
)

type AppVersion struct {
	gorm.Model
	IDService   int32     `gorm:"column:id_service;comment:관련 서비스의 ID"`
	PackageName string    `gorm:"column:package_name;type:varchar(255);not null;comment:앱 패키지명" json:"packageName"`
	VersionName string    `gorm:"column:version_name;type:varchar(255);not null;comment:버전명(예: 1.0.0)" json:"versionName"`
	Type        string    `gorm:"column:type;type:varchar(3);not null;check:type IN ('ios','aos');default:'aos';uniqueIndex:idx_version_type;comment:앱 플랫폼 타입(ios/aos)" json:"type"`
	VersionCode int32     `gorm:"column:version_code;not null;uniqueIndex:idx_version_type;comment:버전 코드(정수값)" json:"versionCode"`
	Description string    `gorm:"column:description;type:varchar(255);comment:버전 설명" json:"description"`
	IsForce     bool      `gorm:"column:is_force;not null;default:false;comment:강제 업데이트 여부" json:"isForce"`
	IsActive    bool      `gorm:"column:is_active;not null;default:false;comment:활성화 상태 여부" json:"isActive"`
	ApplyAt     time.Time `gorm:"column:apply_at;type:timestamptz;comment:버전 적용 시간" json:"applyAt"`
	EffectiveAt time.Time `gorm:"column:effective_at;type:timestamptz;comment:버전 적용 시작 시간" json:"effectiveAt"`
}

func (AppVersion) TableName() string {
	return "app_versions"
}
