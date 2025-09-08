package dbmanager

import (
	"gorm.io/gorm"
)

// AdminRolePermissions 관리자 역할별 권한 설정 테이블
type AdminRolePermissions struct {
	gorm.Model
	// RoleType 관리자 역할 유형 (1~10 범위)
	RoleType int8 `gorm:"column:role_type;not null;comment:'관리자 역할 유형 (1~10 범위)'" json:"roleType"`
	// RoleName 역할명
	RoleName string `gorm:"column:role_name;type:varchar(100);not null;comment:'역할명'" json:"roleName"`

	// 채널 설정 권한들
	ChannelCreate      bool `gorm:"column:channel_create;default:false;comment:'채널 생성 권한'" json:"channelCreate"`
	ChannelModify      bool `gorm:"column:channel_modify;default:false;comment:'채널 수정 권한'" json:"channelModify"`
	ChannelDelete      bool `gorm:"column:channel_delete;default:false;comment:'채널 삭제 권한'" json:"channelDelete"`
	ChannelGroupManage bool `gorm:"column:channel_group_manage;default:false;comment:'채널 그룹 관리 권한'" json:"channelGroupManage"`

	// DB 관리 권한들
	DbBackup          bool `gorm:"column:db_backup;default:false;comment:'DB 백업 권한'" json:"dbBackup"`
	DbDelete          bool `gorm:"column:db_delete;default:false;comment:'DB 삭제 권한'" json:"dbDelete"`
	DuplicateDbManage bool `gorm:"column:duplicate_db_manage;default:false;comment:'중복 DB 관리 권한'" json:"duplicateDbManage"`

	// 계정 관리 권한들
	AccountCreate bool `gorm:"column:account_create;default:false;comment:'계정 생성 권한'" json:"accountCreate"`
	AccountModify bool `gorm:"column:account_modify;default:false;comment:'계정 수정 권한'" json:"accountModify"`
	AccountDelete bool `gorm:"column:account_delete;default:false;comment:'계정 삭제 권한'" json:"accountDelete"`

	// DB 배정 관리 권한들
	BackupOrder bool `gorm:"column:backup_order;default:false;comment:'백업 순서 배정 권한'" json:"backupOrder"`
	TeamBackup  bool `gorm:"column:team_backup;default:false;comment:'팀 백업 권한'" json:"teamBackup"`
	DataBackup  bool `gorm:"column:data_backup;default:false;comment:'데이터 백업 권한'" json:"dataBackup"`

	// 사이트 설정 권한들
	AdChannelManage   bool `gorm:"column:ad_channel_manage;default:false;comment:'광고 채널 관리 권한'" json:"adChannelManage"`
	ApplicationManage bool `gorm:"column:application_manage;default:false;comment:'신청서 관리 권한'" json:"applicationManage"`
}

// TableName 테이블명 반환
func (AdminRolePermissions) TableName() string {
	return "admin_role_permissions"
}

// UserSpecificPermissions 사용자별 개별 권한 설정 테이블
type UserSpecificPermissions struct {
	gorm.Model
	// AdminMemberID 관리자 회원 ID
	IDAdminMember uint `gorm:"column:id_admin_member;not null;comment:'관리자 회원 ID'" json:"IdAdminMember"`
	// RoleType 관리자 역할 유형 (1~10 범위)
	RoleType int8 `gorm:"column:role_type;not null;comment:'관리자 역할 유형 (1~10 범위)'" json:"roleType"`
	// RoleName 역할명
	RoleName string `gorm:"column:role_name;type:varchar(100);not null;comment:'역할명'" json:"roleName"`

	// 채널 설정 권한들
	ChannelCreate      bool `gorm:"column:channel_create;default:false;comment:'채널 생성 권한'" json:"channelCreate"`
	ChannelModify      bool `gorm:"column:channel_modify;default:false;comment:'채널 수정 권한'" json:"channelModify"`
	ChannelDelete      bool `gorm:"column:channel_delete;default:false;comment:'채널 삭제 권한'" json:"channelDelete"`
	ChannelGroupManage bool `gorm:"column:channel_group_manage;default:false;comment:'채널 그룹 관리 권한'" json:"channelGroupManage"`

	// DB 관리 권한들
	DbBackup          bool `gorm:"column:db_backup;default:false;comment:'DB 백업 권한'" json:"dbBackup"`
	DbDelete          bool `gorm:"column:db_delete;default:false;comment:'DB 삭제 권한'" json:"dbDelete"`
	DuplicateDbManage bool `gorm:"column:duplicate_db_manage;default:false;comment:'중복 DB 관리 권한'" json:"duplicateDbManage"`

	// 계정 관리 권한들
	AccountCreate bool `gorm:"column:account_create;default:false;comment:'계정 생성 권한'" json:"accountCreate"`
	AccountModify bool `gorm:"column:account_modify;default:false;comment:'계정 수정 권한'" json:"accountModify"`
	AccountDelete bool `gorm:"column:account_delete;default:false;comment:'계정 삭제 권한'" json:"accountDelete"`

	// DB 배정 관리 권한들
	BackupOrder bool `gorm:"column:backup_order;default:false;comment:'백업 순서 배정 권한'" json:"backupOrder"`
	TeamBackup  bool `gorm:"column:team_backup;default:false;comment:'팀 백업 권한'" json:"teamBackup"`
	DataBackup  bool `gorm:"column:data_backup;default:false;comment:'데이터 백업 권한'" json:"dataBackup"`

	// 사이트 설정 권한들
	AdChannelManage   bool `gorm:"column:ad_channel_manage;default:false;comment:'광고 채널 관리 권한'" json:"adChannelManage"`
	ApplicationManage bool `gorm:"column:application_manage;default:false;comment:'신청서 관리 권한'" json:"applicationManage"`
}

// TableName 테이블명 반환
func (UserSpecificPermissions) TableName() string {
	return "user_specific_permissions"
}

// MenuCategory 메뉴 권한 분류 테이블
type MenuCategory struct {
	gorm.Model
	Name     string `gorm:"column:name;type:varchar(100);not null" json:"name" comment:"메뉴명"`
	Desc     string `gorm:"column:desc;type:varchar(255);default:''" json:"desc" comment:"설명"`
	ActiveYn bool   `gorm:"column:active_yn;type:boolean;default:false" json:"active_yn" comment:"활성여부"`
}

// TableName 테이블명 반환
func (MenuCategory) TableName() string {
	return "menu_categories"
}

// MenuPermission 메뉴 권한 테이블
type MenuPermission struct {
	gorm.Model
	IDMenuCategory uint   `gorm:"column:id_menu_category;not null" json:"idMenuCategory" comment:"메뉴분류인덱스"`
	Name           string `gorm:"column:name;type:varchar(100);not null" json:"name" comment:"권한명"`
	Desc           string `gorm:"column:desc;type:varchar(255);default:''" json:"desc" comment:"설명"`
	ActiveYn       bool   `gorm:"column:active_yn;type:boolean;default:false" json:"active_yn" comment:"활성여부"`
}

// TableName 테이블명 반환
func (MenuPermission) TableName() string {
	return "menu_permissions"
}

// UserMenuPermission 사용자 메뉴 권한 테이블
type UserMenuPermission struct {
	gorm.Model
	ID               uint `gorm:"column:id;primaryKey" json:"id" comment:"사용자 권한 인덱스"`
	IDMenuPermission uint `gorm:"column:id_munu_permission;not null" json:"idMenuPermission" comment:"메뉴 인덱스"`
	IDAdminMember    uint `gorm:"column:id_admin_member;not null" json:"idAdminMember" comment:"관리자 인덱스"`
}

// TableName 테이블명 반환
func (UserMenuPermission) TableName() string {
	return "user_menu_permissions"
}
