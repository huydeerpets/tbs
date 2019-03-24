package movie

import (
	"github.com/huydeerpets/tbs/tests"

	. "gopkg.in/check.v1"
)

type TestYoutube struct {
	test.TSuite
}

func init() {
	var t test.Accessor = &TestYoutube{}
	t.SetTableNameList([]string{})

	var _ = Suite(t)
}
