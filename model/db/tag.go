package db

type Tag struct {
	Base
	AppID string `gorm:"not null;index"`
	Name  string `gorm:"not null;index"`
	Color string `grom:"not null"`
}
