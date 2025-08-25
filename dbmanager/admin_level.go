package dbmanager

import (
	"gorm.io/gorm"
)

// AdminLevel 관리자 등급별 권한 설정 구조체
type AdminLevel struct {
	gorm.Model
	// UserLevel 사용자 등급 (1:최고관리자, 2:일반관리자, 3:팀관리자, 4:팀원)
	UserLevel int8 `gorm:"column:level;not null;type:default:4;comment:'등급(1:최고관리자,2:일반관리자,3:팀관리자,4:팀원)'" json:"userLevel"`
	// PermissionAll 전체/원본 데이터 접근 권한
	PermissionAll bool `gorm:"column:permission_all;default:false;comment:'전체/원본 권한'" json:"permissionAll"`
	// PermissionDupDb 중복 데이터베이스 관리 권한
	PermissionDupDb bool `gorm:"column:permission_dup_db;default:false;comment:'중복DB 권한'" json:"permissionDupDb"`
	// PermissionTeamAssign 팀 배정 관리 권한
	PermissionTeamAssign bool `gorm:"column:permission_team_assign;default:false;comment:'팀 배정 권한'" json:"permissionTeamAssign"`
	// PermissionMemberAssign 팀원 배정 관리 권한
	PermissionMemberAssign bool `gorm:"column:permission_member_assign;default:false;comment:'팀원 배정 권한'" json:"permissionMemberAssign"`
	// PermissionDelete 데이터 삭제 권한
	PermissionDelete bool `gorm:"column:permission_delete;default:false;comment:'삭제 권한'" json:"permissionDelete"`
	// PermissionAdChannel 광고 채널 관리 권한
	PermissionAdChannel bool `gorm:"column:permission_ad_channel;default:false;comment:'광고 채널 권한'" json:"permissionAdChannel"`
	// PermissionTeamAssignByTeam 팀별 배정 관리 권한
	PermissionTeamAssignByTeam bool `gorm:"column:permission_team_assign_by_team;default:false;comment:'팀별 배정 권한'" json:"permissionTeamAssignByTeam"`
	// PermissionAdConcept 광고 컨셉 관리 권한
	PermissionAdConcept bool `gorm:"column:permission_ad_concept;default:false;comment:'광고 컨셉 권한'" json:"permissionAdConcept"`
	// PermissionApplication 신청서 관리 권한
	PermissionApplication bool `gorm:"column:permission_application;default:false;comment:'신청서 권한'" json:"permissionApplication"`
	// PermissionBackup 백업 관리 권한
	PermissionBackup bool `gorm:"column:permission_backup;default:false;comment:'백업 권한'" json:"permissionBackup"`
	// PermissionTeamAssignSecond 두번째 팀 배정 관리 권한
	PermissionTeamAssignSecond bool `gorm:"column:permission_team_assign_second;default:false;comment:'팀 배정(두번째) 권한'" json:"permissionTeamAssignSecond"`
	// PermissionMemberAssignSecond 두번째 팀원 배정 관리 권한
	PermissionMemberAssignSecond bool `gorm:"column:permission_member_assign_second;default:false;comment:'팀원 배정(두번째) 권한'" json:"permissionMemberAssignSecond"`
	// PermissionPersonal 개인 데이터 접근 권한
	PermissionPersonal bool `gorm:"column:permission_personal;default:false;comment:'개인 권한'" json:"permissionPersonal"`
}

// TableName AdminLevel 테이블명 반환
func (AdminLevel) TableName() string {
	return "admin_levels"
}
