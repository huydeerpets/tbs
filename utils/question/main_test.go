package question

import (
	"github.com/huydeerpets/tbs/tests"

	. "gopkg.in/check.v1"
)

type TestMain struct {
	test.TSuite
}

func init() {
	var t test.Accessor = &TestMain{}
	t.SetTableNameList([]string{
		"log_questions",
	})

	var _ = Suite(t)
}

func (t *TestMain) TestAdd(c *C) {
	r := Add(1, "abc", "abc@test.com")

	c.Check(r, Equals, nil)
}