package controllersFollow

import (
	"net/http"
	"net/http/httptest"
	"testing"

	_ "github.com/huydeerpets/tbs/routers"
	"github.com/huydeerpets/tbs/tests"

	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
)

func setupDelete() {
	test.Setup()
	test.SetupFixture([]string{
		"user_masters",
		"user_contributions",
		"user_contribution_follows",
	})
}

func TestDelete(t *testing.T) {
	setupDelete()

	r, err := http.NewRequest(
		"DELETE",
		"/api/follows/1/?user_id=1000",
		nil,
	)

	if err != nil {
		panic(err)
	}

	r.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	Convey("/follows/1\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}
