package models

// LogContributionImagePostImageLog
type LogContributionImage struct {
	BaseModel
	UserContributionID int `json:"user_contribution_id"`
}

// GetIDAndAdd Add and get ID
func (l *LogContributionImage) GetIDAndAdd() (uint, error) {
	if err := Create(l); err != nil {
		return 0, err
	}

	return l.ID, nil
}
