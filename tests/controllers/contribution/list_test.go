package controllersContribution

import (
	"net/http"
	"net/http/httptest"
	"testing"

	_ "github.com/huydeerpets/tbs/routers"
	"github.com/huydeerpets/tbs/tests"

	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
)

func setUpList() {
	test.Setup()
	test.SetupFixture([]string{
		"user_masters",
		"user_contributions",
		"user_contribution_details",
		"user_contribution_tags",
	})
}

func TestListPost(t *testing.T) {
	setUpList()

	r, err := http.NewRequest(
		"GET",
		"/api/contributions/list/2",
		nil,
	)
	if err != nil {
		panic(err)
	}

	r.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	Convey("/contribution/list/\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}
