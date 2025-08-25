package dbmanager

import (
	"gorm.io/gorm"
)

type AdminLevel struct {
	gorm.Model
	Level                        int16  `gorm:"column:level;not null;type:smallint;comment:'등급(1:최고관리자,2:일반관리자,3:팀관리자,4:팀원)'" json:"level"`
	PermissionAll                bool   `gorm:"column:permission_all;default:false;comment:'전체/원본 권한'" json:"permissionAll"`
	PermissionDupDb              bool   `gorm:"column:permission_dup_db;default:false;comment:'중복DB 권한'" json:"permissionDupDb"`
	PermissionTeamAssign         bool   `gorm:"column:permission_team_assign;default:false;comment:'팀 배정 권한'" json:"permissionTeamAssign"`
	PermissionMemberAssign       bool   `gorm:"column:permission_member_assign;default:false;comment:'팀원 배정 권한'" json:"permissionMemberAssign"`
	PermissionDelete             bool   `gorm:"column:permission_delete;default:false;comment:'삭제 권한'" json:"permissionDelete"`
	PermissionAdChannel          bool   `gorm:"column:permission_ad_channel;default:false;comment:'광고 채널 권한'" json:"permissionAdChannel"`
	PermissionTeamAssignByTeam   bool   `gorm:"column:permission_team_assign_by_team;default:false;comment:'팀별 배정 권한'" json:"permissionTeamAssignByTeam"`
	PermissionAdConcept          bool   `gorm:"column:permission_ad_concept;default:false;comment:'광고 컨셉 권한'" json:"permissionAdConcept"`
	PermissionApplication        bool   `gorm:"column:permission_application;default:false;comment:'신청서 권한'" json:"permissionApplication"`
	PermissionBackup             bool   `gorm:"column:permission_backup;default:false;comment:'백업 권한'" json:"permissionBackup"`
	PermissionTeamAssignSecond   bool   `gorm:"column:permission_team_assign_second;default:false;comment:'팀 배정(두번째) 권한'" json:"permissionTeamAssignSecond"`
	PermissionMemberAssignSecond bool   `gorm:"column:permission_member_assign_second;default:false;comment:'팀원 배정(두번째) 권한'" json:"permissionMemberAssignSecond"`
	PermissionPersonal           bool   `gorm:"column:permission_personal;default:false;comment:'개인 권한'" json:"permissionPersonal"`
}

func (AdminLevel) TableName() string {
	return "admin_levels"
}
