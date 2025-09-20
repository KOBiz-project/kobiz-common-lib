package dbmanager

// MenuCategory 메뉴 권한 분류 테이블
type MenuCategory struct {
	ModelCommon
	Name string `gorm:"column:name;type:varchar(100);not null" json:"name" comment:"메뉴명"`
}

// TableName 테이블명 반환
func (MenuCategory) TableName() string {
	return "menu_categories"
}

// MenuPermission 메뉴 권한 테이블
type MenuPermission struct {
	ModelCommon
	IDMenuCategory uint   `gorm:"column:id_menu_category;not null" json:"idMenuCategory" comment:"메뉴분류인덱스"`
	Name           string `gorm:"column:name;type:varchar(100);not null" json:"name" comment:"권한명"`
	Code           string `gorm:"column:code;type:varchar(100);not null;unique" json:"code" comment:"권한코드"`
}

// TableName 테이블명 반환
func (MenuPermission) TableName() string {
	return "menu_permissions"
}

// LevelPermission 레벨별 권한 테이블
type LevelPermission struct {
	ModelCommon
	Name     string `gorm:"column:name;type:varchar(100);not null" json:"name" comment:"레벨별 권한 명"`
	RoleType int8   `gorm:"column:role_type;type:tinyint;default:1;not null" json:"roleType" comment:"관리자 역할 유형 (1~10 범위)"`
}

// TableName 테이블명 반환
func (LevelPermission) TableName() string {
	return "level_permissions"
}

// LevelPermissionsMenuPermissions 레벨별 권한과 메뉴 권한 매핑 테이블
type LevelPermissionsMenuPermissions struct {
	ModelCommon
	IDLevelPermission uint `gorm:"column:id_level_permission;not null" json:"idLevelPermission" comment:"레벨별 권한 인덱스"`
	IDMenuPermission  uint `gorm:"column:id_menu_permission;not null" json:"idMenuPermission" comment:"메뉴 권한 인덱스"`
}

// TableName 테이블명 반환
func (LevelPermissionsMenuPermissions) TableName() string {
	return "level_permissions_menu_permissions"
}
