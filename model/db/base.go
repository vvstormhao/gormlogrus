package db

import (
	"time"
)

// Base 数据库表基础字段
type Base struct {
	ID        uint      `gorm:"primary_key"`
	CreatedAt time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

// SoftDelBase 支持伪删除
type SoftDelBase struct {
	Base
	DeletedAt *time.Time
}
