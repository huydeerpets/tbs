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
		"user_contribution_follows",
		"contribution_total_follows",
		"user_contribution_searches",
	})
}

func TestExec(t *testing.T) {
	r := AddContributionTotalFollows()

	Convey("tasks/contributionTotalFollows/main.go\n", t, func() {
		Convey("AddContributionTotalFollows", func() {
			So(r, ShouldEqual, nil)
		})
	})

	r = SaveUserContributionSearchToFollowCount()

	Convey("tasks/contributionTotalFollows/main.go\n", t, func() {
		Convey("SaveUserContributionSearchToFollowCount", func() {
			So(r, ShouldEqual, nil)
		})
	})

}
