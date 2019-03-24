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
		"user_contribution_sound_details",
		"user_contribution_movies",
	})
}

func TestExec(t *testing.T) {
	Convey("tasks/removeContribution/main.go\n", t, func() {
		r := RemoveSoundDetail()

		Convey("RemoveSoundDetail", func() {
			So(r, ShouldEqual, nil)
		})

		r = RemoveJoinFile()

		Convey("RemoveJoinFile", func() {
			So(r, ShouldEqual, nil)
		})
	})
}
