package models

// LogBugReport Bug Report
type LogBugReport struct {
	BaseModel
	UserID int    `json:"user_id"`
	Body   string `json:"body"`
}

// Add Add
func (c *LogBugReport) Add() error {
	return Create(c)
}
