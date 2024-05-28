package models

type Conversation struct {
	ID           uint `json:"id"`
	Participant1 uint `json:"participant1"`
	Participant2 uint `json:"participant2"`
}
