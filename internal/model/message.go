package models

import "time"

type Message struct {
	ID             uint      `json:"id" gorm:"primaryKey"`
	ConversationID uint      `json:"conversation_id"`
	SenderID       uint      `json:"sender_id"`
	Content        string    `json:"content"`
	SendedAt       time.Time `json:"sended_at"`
}
