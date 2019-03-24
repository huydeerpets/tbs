package routers

import (
	"github.com/huydeerpets/tbs/controllers"
	"github.com/huydeerpets/tbs/controllers/bug"
	"github.com/huydeerpets/tbs/controllers/character"
	"github.com/huydeerpets/tbs/controllers/contribution"
	"github.com/huydeerpets/tbs/controllers/facebook"
	"github.com/huydeerpets/tbs/controllers/follow"
	"github.com/huydeerpets/tbs/controllers/forget_password"
	"github.com/huydeerpets/tbs/controllers/google"
	"github.com/huydeerpets/tbs/controllers/login"
	"github.com/huydeerpets/tbs/controllers/movie"
	"github.com/huydeerpets/tbs/controllers/native"
	"github.com/huydeerpets/tbs/controllers/problem"
	"github.com/huydeerpets/tbs/controllers/profile"
	"github.com/huydeerpets/tbs/controllers/question"
	"github.com/huydeerpets/tbs/controllers/sound"
	"github.com/huydeerpets/tbs/controllers/tag"
	"github.com/huydeerpets/tbs/controllers/twitter"
	"github.com/huydeerpets/tbs/controllers/user"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/api/bug/", &controllersBug.AddController{})

	beego.Router("/api/contributions/list/:order([0-9]+)", &controllersContribution.ListController{})
	beego.Router("/api/contributions/new/", &controllersContribution.NewController{})
	beego.Router("/api/contributions/:id([0-9]+)", &controllersContribution.MainController{})
	beego.Router("/api/contributions/upload/", &controllersContribution.UploadController{})
	beego.Router("/api/contributions/edit/:id([0-9]+)", &controllersContribution.EditController{})
	beego.Router("/api/contributions/search/", &controllersContribution.SearchController{})

	beego.Router("/api/characters/", &controllersCharacter.MainController{})
	beego.Router("/api/characters/:id([0-9]+)", &controllersCharacter.DeleteController{})

	beego.Router("/api/facebook/oauth/", &controllersFacebook.OauthController{})
	beego.Router("/api/facebook/callback/", &controllersFacebook.CallbackController{})

	beego.Router("/api/follows/:id([0-9]+)", &controllersFollow.MainController{})
	beego.Router("/api/follows/list/", &controllersFollow.ListController{})

	beego.Router("/api/forget_password/", &controllersForgetPassword.MainController{})
	beego.Router("/api/forget_password/check/:email/:keyword", &controllersForgetPassword.CheckController{})

	beego.Router("/api/login/auth/", &controllersLogin.AuthController{})
	beego.Router("/api/login/check/", &controllersLogin.CheckController{})
	beego.Router("/api/login/callback/", &controllersLogin.CallbackController{})

	beego.Router("/api/logout/", &controllersLogin.LogoutController{})

	beego.Router("/api/google/oauth/", &controllersGoogle.OauthController{})
	beego.Router("/api/google/callback/", &controllersGoogle.CallbackController{})

	beego.Router("/api/movies/:id([0-9]+)", &controllersMovie.MainController{})
	beego.Router("/api/movies/:id([0-9]+)/upload/", &controllersMovie.UploadController{})
	beego.Router("/api/movies/connect/:id([0-9]+)", &controllersMovie.ConnectController{})
	beego.Router("/api/movies/callback/", &controllersMovie.CallbackController{})

	beego.Router("/api/native/callback/", &controllersNative.CallbackController{})
	beego.Router("/api/native/dev-callback/", &controllersNative.DevCallbackController{})
	beego.Router("/api/native/redirect/", &controllersNative.RedirectController{})
	beego.Router("/api/native/dev-redirect/", &controllersNative.DevRedirectController{})

	beego.Router("/api/problem/", &controllersProblem.AddController{})

	beego.Router("/api/profile/", &controllersUserProfile.UploadController{})

	beego.Router("/api/question/", &controllersQuestion.AddController{})

	beego.Router("/api/sounds/:id([0-9]+)/", &controllersSound.MainController{})
	beego.Router("/api/sounds/:id([0-9]+)/make/", &controllersSound.MakeController{})
	beego.Router("/api/sounds/:id([0-9]+)/reflect/", &controllersSound.ReflectController{})
	beego.Router("/api/sounds/:id([0-9]+)/voice/all/", &controllersSound.SaveVoiceListController{})
	beego.Router("/api/sounds/body/", &controllersSound.SaveBodyController{})
	beego.Router("/api/sounds/voice/", &controllersSound.SaveVoiceController{})
	beego.Router("/api/sounds/length/", &controllersSound.LengthController{})

	beego.Router("/api/tags/", &controllersTag.MainController{})

	beego.Router("/api/twitter/oauth/", &controllersTwitter.OauthController{})
	beego.Router("/api/twitter/callback/", &controllersTwitter.CallbackController{})

	beego.Router("/api/me/", &controllersUser.MainController{})
	beego.Router("/api/users/new/", &controllersLogin.NewController{})
	beego.Router("/api/users/contribution/list/", &controllersUser.ContributionListController{})

	beego.Router("/*", &controllers.MainController{})
}
