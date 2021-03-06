package twitter

import (
	"github.com/huydeerpets/tbs/tests"

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

func (t *TestMain) TestGetConnect(c *C) {
	r := GetConnect()

	c.Check(r, Not(Equals), nil)
}
