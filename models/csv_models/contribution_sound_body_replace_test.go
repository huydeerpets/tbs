package csvModels

import (
	"github.com/huydeerpets/tbs/tests"

	. "gopkg.in/check.v1"
)

type TestContributionSoundBodyReplace struct {
	test.TSuite
}

func init() {
	var t test.Accessor = &TestContributionSoundBodyReplace{}
	t.SetTableNameList([]string{})

	var _ = Suite(t)
}

func (t *TestContributionSoundBodyReplace) TestGetStructAll(c *C) {

	m := &ContributionSoundBodyReplace{}

	r, _ := m.GetStructAll()

	c.Check(r[0].ID, Equals, "1")
}
