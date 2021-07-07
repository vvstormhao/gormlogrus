package db

type AppKey struct {
	Base
	AppKey    string `gorm:"not null;unique_index" `
	AppSecret string `gorm:"not null" `
}
