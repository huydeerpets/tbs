package models

const (
	// ProblemTypeSpam Problem type: Spam
	ProblemTypeSpam = 1
	// ProblemTypeInappropriate Problem type: Inappropriate
	ProblemTypeInappropriate = 2
)

// LogProblemContributionPostProblemReport
type LogProblemContributionReport struct {
	BaseModel
	UserID             int `json:"user_id"`
	Type               int `json:"type"`
	UserContributionID int `json:"user_contribution_id"`
}

// Add Add
func (l *LogProblemContributionReport) Add() error {
	return Create(l)
}
