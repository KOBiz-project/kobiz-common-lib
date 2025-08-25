package dbmanager

import "gorm.io/gorm"

// AdminTeam 관리자 팀 구조체
type AdminTeam struct {
	gorm.Model
	// TeamName 팀 이름
	TeamName string `gorm:"column:team_name;not null;type:varchar(255);comment:'팀 이름'" json:"teamName"`
	// activeYn 활성화 여부
	ActiveYn bool `gorm:"column:active_yn;default:false;comment:'활성화 여부'" json:"activeYn"`
}

// TableName AdminTeam 테이블명 반환
func (AdminTeam) TableName() string {
	return "admin_teams"
}
