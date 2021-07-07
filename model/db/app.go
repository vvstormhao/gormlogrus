package db

// App apps表结构
type App struct {
	SoftDelBase
	AppID string `gorm:"not null;unique_index" ` // 公司ID
}
