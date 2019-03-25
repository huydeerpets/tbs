package csvModels

// ContributionSoundBodyReplace Voice text replacement
type ContributionSoundBodyReplace struct {
	ID      string
	Text    string
	Replace string
	Delete  string
}

// GetStructAll Get everything
func (c *ContributionSoundBodyReplace) GetStructAll() (r []ContributionSoundBodyReplace, err error) {
	err = GetAll("contribution_sound_body_replace.csv", &r)

	return r, err
}
