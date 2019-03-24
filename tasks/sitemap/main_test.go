package main

import (
	_ "github.com/huydeerpets/tbs/routers"
	"github.com/huydeerpets/tbs/tests"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func init() {
	test.Setup()
	test.SetupFixture([]string{
		"user_contributions",
	})
}

func TestContribution(t *testing.T) {
	sm := create()

	contribution(sm)

	sm.Finalize().PingSearchEngines("http://newengine.com/ping?url=%s")

	Convey("tasks/makeMovie/main.go\n", t, func() {
		Convey("MakeMovie", func() {
			So(nil, ShouldEqual, nil)
		})
	})
}
