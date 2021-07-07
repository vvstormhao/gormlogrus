package db

import (
	"time"
)

const (
	StatusNonQA    = 0 // 未质检
	StatusDoneQA   = 1 // 质检完成
	StatusJustice  = 2 // 申诉
	StatusFeedback = 3 // 驳回
)

type Agent struct {
	ID   string
	Name string
}

type Group struct {
	ID   string
	Name string
}

type Task struct {
	Base
	AppID          string    `gorm:"not null;index"`
	SessionID      string    `gorm:"not null;index"`
	Channel        string    `gorm:"not null;index"`
	StartedAt      time.Time `gorm:"not null"`
	EndedAt        time.Time `gorm:"not null"`
	Duration       int       `gorm:"not null"`
	Caller         string
	AgentID        string
	AgentName      string
	QaID           string
	QaName         string
	AgentGroupID   string
	AgentGroupName string
	Suvery         int
	RecordURL      string
	DialogDetail   string `gorm:"type:TEXT;not null"`
	Score          int
	Tags           string
	Comment        string `gorm:"type:TEXT"`
	Status         int    `gorm:"not null;default:0"`
}
