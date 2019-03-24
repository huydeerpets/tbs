package controllersUser

import (
	"net/http"
	"net/http/httptest"
	"testing"

	_ "github.com/huydeerpets/tbs/routers"
	"github.com/huydeerpets/tbs/tests"

	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
)

func setUpUpload() {
	test.Setup()

	test.SetupFixture([]string{
		"user_profile_images",
	})
}

func TestUpload(t *testing.T) {
	setUpUpload()

	r, err := http.NewRequest(
		"POST",
		"/api/profile/",
		nil,
	)

	if err != nil {
		panic(err)
	}

	r.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	Convey("/profile/\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}
