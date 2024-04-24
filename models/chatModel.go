package models

import (
	"gorm.io/gorm"
	"time"
)

type Chat struct {
	gorm.Model
	ID       uint      `gorm:"primaryKey"`
	Room     string    `gorm:"not null"`
	UserOne  string    `gorm:"not null"`
	UserTwo  string    `gorm:"not null"`
	Read     bool      `gorm:"default:false"`
	Messages []Message `gorm:"foreignKey:ChatID"`
}

type Message struct {
	gorm.Model
	ChatID    uint      `gorm:"not null"`
	Chat      Chat      `gorm:"foreignKey:ChatID;references:ID"`
	Content   string    `gorm:"not null"`
	UserID    uint      `gorm:"not null"`
	User      User      `gorm:"foreignKey:UserID;references:ID"`
	Timestamp time.Time `gorm:"not null;autoCreateTime"` // autoCreateTime tag to automatically set the time
}

// BeforeCreate hook to set Timestamp to current time before creating a Message
func (m *Message) BeforeCreate(tx *gorm.DB) error {
	m.Timestamp = time.Now()
	return nil
}
