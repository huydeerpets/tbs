package models

// LogUserContributionPostLog
type LogUserContribution struct {
	BaseModel
	UserContributionID int `json:"user_contribution_id"`
	UserID             int `json:"user_id"`
}

// Add Add
func (l *LogUserContribution) Add() error {
	return Create(l)
}
