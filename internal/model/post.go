package models

type Post struct {
	Subject     string `json:"subject"`
	UserID      uint   `json:"user_id"`
	Description string `json:"description"`
}
