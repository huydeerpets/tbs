package movie

import (
	"github.com/huydeerpets/tbs/models/csv_models"
	"github.com/huydeerpets/tbs/tests"
	"github.com/huydeerpets/tbs/utils/sound"

	. "gopkg.in/check.v1"
)

type TestMain struct {
	test.TSuite
}

func init() {
	var t test.Accessor = &TestMain{}
	t.SetTableNameList([]string{})

	var _ = Suite(t)
}

func (t *TestMain) TestExecMakeMovie(c *C) {
	r := ExecMakeMovie(0)

	c.Check(r, Equals, nil)
}

func (t *TestMain) TestRemoveFile(c *C) {
	list := []string{
		"0_1",
		"0_2",
	}

	sound.AddTmpSound("こんにちは", list[0], csvModels.VoiceTypeMeiNormal)
	sound.AddTmpSound("おはようございます", list[1], csvModels.VoiceTypeMeiNormal)
	sound.Join(list, "0")

	sound.ToM4a("0")

	r := Make("0")
	c.Check(r, Equals, nil)

	r = ToFilter("0")
	c.Check(r, Equals, nil)

	r = RemoveFile("0")

	c.Check(r, Equals, nil)
}
