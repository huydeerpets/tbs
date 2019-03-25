package models

// LogQuestion QuestionLog
type LogQuestion struct {
	BaseModel
	UserID int    `json:"user_id"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}

// Add Add
func (l *LogQuestion) Add() error {
	return Create(l)
}
