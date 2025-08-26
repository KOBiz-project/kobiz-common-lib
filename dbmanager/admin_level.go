package dbmanager

import (
	"gorm.io/gorm"
)

// AdminLevel 관리자 등급 전체 권한 설정 구조체
type AdminLevel struct {
	gorm.Model
	// UserLevel 사용자 등급 (1:최고관리자, 2:일반관리자, 3:팀관리자, 4:팀원)
	UserLevel int8 `gorm:"column:level;not null;default:4;comment:'사용자 등급 (1:최고관리자, 2:일반관리자, 3:팀관리자, 4:팀원)'" json:"userLevel"`
	// PermissionAll 전체/원본 데이터 접근 권한 (DB 확인)
	PermissionAll bool `gorm:"column:permission_all;default:false;comment:'전체/원본 데이터 접근 권한 (DB 확인)'" json:"permissionAll"`
	// PermissionDupDb 중복 데이터베이스 관리 권한 (DB 확인)
	PermissionDupDb bool `gorm:"column:permission_dup_db;default:false;comment:'중복 데이터베이스 관리 권한 (DB 확인)'" json:"permissionDupDb"`
	// PermissionTeamAssign 팀 배정 관리 권한 (DB 확인)
	PermissionTeamAssign bool `gorm:"column:permission_team_assign;default:false;comment:'팀 배정 관리 권한 (DB 확인)'" json:"permissionTeamAssign"`
	// PermissionMemberAssign 팀원 배정 관리 권한 (DB 확인)
	PermissionMemberAssign bool `gorm:"column:permission_member_assign;default:false;comment:'팀원 배정 관리 권한 (DB 확인)'" json:"permissionMemberAssign"`
	// PermissionDelete 데이터 삭제 권한 (DB 관리)
	PermissionDelete bool `gorm:"column:permission_delete;default:false;comment:'데이터 삭제 권한 (DB 관리)'" json:"permissionDelete"`
	// PermissionBackup 백업 관리 권한 (DB 관리)
	PermissionBackup bool `gorm:"column:permission_backup;default:false;comment:'백업 관리 권한 (DB 관리)'" json:"permissionBackup"`
	// PermissionTeamAssignSecond 팀 배정 관리 권한 (DB 배정)
	PermissionTeamAssignSecond bool `gorm:"column:permission_team_assign_second;default:false;comment:'팀 배정 관리 권한 (DB 배정)'" json:"permissionTeamAssignSecond"`
	// PermissionMemberAssignSecond 팀원 배정 관리 권한 (DB 배정)
	PermissionMemberAssignSecond bool `gorm:"column:permission_member_assign_second;default:false;comment:'팀원 배정 관리 권한 (DB 배정)'" json:"permissionMemberAssignSecond"`
	// PermissionTeam 팀 데이터 접근 권한 (계정 관리)
	PermissionTeam bool `gorm:"column:permission_team;default:false;comment:'팀 데이터 접근 권한 (계정 관리)'" json:"permissionTeam"`
	// PermissionPersonal 개인 데이터 접근 권한 (계정 관리)
	PermissionPersonal bool `gorm:"column:permission_personal;default:false;comment:'개인 데이터 접근 권한 (계정 관리)'" json:"permissionPersonal"`
	// PermissionAdChannel 광고 채널 관리 권한 (설정)
	PermissionAdChannel bool `gorm:"column:permission_ad_channel;default:false;comment:'광고 채널 관리 권한 (설정)'" json:"permissionAdChannel"`
	// PermissionTeamAssignByTeam 팀별 배정 관리 권한 (설정)
	PermissionTeamAssignByTeam bool `gorm:"column:permission_team_assign_by_team;default:false;comment:'팀별 배정 관리 권한 (설정)'" json:"permissionTeamAssignByTeam"`
	// PermissionAdConcept 광고 컨셉 관리 권한 (설정)
	PermissionAdConcept bool `gorm:"column:permission_ad_concept;default:false;comment:'광고 컨셉 관리 권한 (설정)'" json:"permissionAdConcept"`
	// PermissionApplication 신청서 관리 권한 (설정)
	PermissionApplication bool `gorm:"column:permission_application;default:false;comment:'신청서 관리 권한 (설정)'" json:"permissionApplication"`
}

// TableName AdminLevel 테이블명 반환
func (AdminLevel) TableName() string {
	return "admin_levels"
}

// AdminLevelUser 관리자 등급 개별 사용자 권한 설정 구조체
type AdminLevelUser struct {
	gorm.Model
	// IdAdminMembers 관리자 회원 ID
	IdAdminMembers uint `gorm:"column:id_admin_members;not null;comment:'관리자 회원 ID'" json:"idAdminMembers"`

	// UserLevel 사용자 등급 (1:최고관리자, 2:일반관리자, 3:팀관리자, 4:팀원)
	UserLevel int8 `gorm:"column:level;not null;default:4;comment:'사용자 등급 (1:최고관리자, 2:일반관리자, 3:팀관리자, 4:팀원)'" json:"userLevel"`
	// PermissionAll 전체/원본 데이터 접근 권한 (DB 확인)
	PermissionAll bool `gorm:"column:permission_all;default:false;comment:'전체/원본 데이터 접근 권한 (DB 확인)'" json:"permissionAll"`
	// PermissionDupDb 중복 데이터베이스 관리 권한 (DB 확인)
	PermissionDupDb bool `gorm:"column:permission_dup_db;default:false;comment:'중복 데이터베이스 관리 권한 (DB 확인)'" json:"permissionDupDb"`
	// PermissionTeamAssign 팀 배정 관리 권한 (DB 확인)
	PermissionTeamAssign bool `gorm:"column:permission_team_assign;default:false;comment:'팀 배정 관리 권한 (DB 확인)'" json:"permissionTeamAssign"`
	// PermissionMemberAssign 팀원 배정 관리 권한 (DB 확인)
	PermissionMemberAssign bool `gorm:"column:permission_member_assign;default:false;comment:'팀원 배정 관리 권한 (DB 확인)'" json:"permissionMemberAssign"`
	// PermissionDelete 데이터 삭제 권한 (DB 관리)
	PermissionDelete bool `gorm:"column:permission_delete;default:false;comment:'데이터 삭제 권한 (DB 관리)'" json:"permissionDelete"`
	// PermissionBackup 백업 관리 권한 (DB 관리)
	PermissionBackup bool `gorm:"column:permission_backup;default:false;comment:'백업 관리 권한 (DB 관리)'" json:"permissionBackup"`
	// PermissionTeamAssignSecond 팀 배정 관리 권한 (DB 배정)
	PermissionTeamAssignSecond bool `gorm:"column:permission_team_assign_second;default:false;comment:'팀 배정 관리 권한 (DB 배정)'" json:"permissionTeamAssignSecond"`
	// PermissionMemberAssignSecond 팀원 배정 관리 권한 (DB 배정)
	PermissionMemberAssignSecond bool `gorm:"column:permission_member_assign_second;default:false;comment:'팀원 배정 관리 권한 (DB 배정)'" json:"permissionMemberAssignSecond"`
	// PermissionTeam 팀 데이터 접근 권한 (계정 관리)
	PermissionTeam bool `gorm:"column:permission_team;default:false;comment:'팀 데이터 접근 권한 (계정 관리)'" json:"permissionTeam"`
	// PermissionPersonal 개인 데이터 접근 권한 (계정 관리)
	PermissionPersonal bool `gorm:"column:permission_personal;default:false;comment:'개인 데이터 접근 권한 (계정 관리)'" json:"permissionPersonal"`
	// PermissionAdChannel 광고 채널 관리 권한 (설정)
	PermissionAdChannel bool `gorm:"column:permission_ad_channel;default:false;comment:'광고 채널 관리 권한 (설정)'" json:"permissionAdChannel"`
	// PermissionTeamAssignByTeam 팀별 배정 관리 권한 (설정)
	PermissionTeamAssignByTeam bool `gorm:"column:permission_team_assign_by_team;default:false;comment:'팀별 배정 관리 권한 (설정)'" json:"permissionTeamAssignByTeam"`
	// PermissionAdConcept 광고 컨셉 관리 권한 (설정)
	PermissionAdConcept bool `gorm:"column:permission_ad_concept;default:false;comment:'광고 컨셉 관리 권한 (설정)'" json:"permissionAdConcept"`
	// PermissionApplication 신청서 관리 권한 (설정)
	PermissionApplication bool `gorm:"column:permission_application;default:false;comment:'신청서 관리 권한 (설정)'" json:"permissionApplication"`
}

// TableName AdminLevel 테이블명 반환
func (AdminLevelUser) TableName() string {
	return "admin_level_users"
}
