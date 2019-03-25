package controllers

import (
	"errors"

	"github.com/huydeerpets/tbs/utils/log"

	"github.com/astaxie/beego"
)

// noUserID NoUserID
const noUserID = 0

// BaseController BaseControler
type BaseController struct {
	beego.Controller
}

// ErrorResponse Errorresponse
type ErrorResponse struct {
	Message string `json:"message"`
	ErrCode int    `json:"errCode"`
}

// Accessor Base accessor
type Accessor interface {
	GetUserID() int
	ServerError()
}

const (
	// ErrCodeCommon Common error
	ErrCodeCommon = 1
	// ErrCodeUserNotFound Unable to get user information or mismatch
	ErrCodeUserNotFound = 2
	// ErrCodeLoginNotFound Not logged in
	ErrCodeLoginNotFound = 3
	// ErrCreateUser User registration failed
	ErrCreateUser = 4
	// ErrParameter Parameter error
	ErrParameter = 5
	// ErrImageConversionImageConversion error
	ErrImageConversion = 6
	// ErrImageResizeImageShrink error
	ErrImageResize = 7
	// ErrContributionNewPostError
	ErrContributionNew = 8
	// ErrContributionSavePostStorage failure
	ErrContributionSave = 9
	// ErrContributionTagSavePostTag save failure
	ErrContributionTagSave = 10
	// ErrUserSave User save failure
	ErrUserSave = 11
	// ErrUserOrPasswordDifferent User or password is different
	ErrUserOrPasswordDifferent = 12
	// ErrContributionSearch Failed to get search
	ErrContributionSearch = 13
	// ErrFollowedFollowAlready
	ErrFollowed = 14
	// ErrAddFollowFollow addition failure
	ErrAddFollow = 15
	// ErrContributionNotFoundPost does not exist
	ErrContributionNotFound = 16
	// ErrDeleteFollowFollowDelete failure
	ErrDeleteFollow = 17
	// The maximum number of ErrTagMaxNumberOver tags is exceeded
	ErrTagMaxNumberOver = 18
	// ErrTagNameOverlap Duplicate tag name exists
	ErrTagNameOverlap = 19
	// Not the user who did ErrContributionNoUserPost
	ErrContributionNoUser = 20
	// ErrPasswordMinLength Password is less than the minimum number of characters
	ErrPasswordMinLength = 21
)

// errResponseMap Error response map
var errResponseMap = map[int]ErrorResponse{
	ErrCodeCommon: {
		Message: "An error has occurred.",
	},
	ErrCodeUserNotFound: {
		Message: "User information could not be obtained. Please login again.",
	},
	ErrCodeLoginNotFound: {
		Message: "This screen can not be used for users who are not logged in",
	},
	ErrCreateUser: {
		Message: "Failed to create user. Please register again.",
	},
	ErrParameter: {
		Message: "Bad parameter has been sent",
	},
	ErrImageConversion: {
		Message: "Image conversion failed",
	},
	ErrImageResize: {
		Message: "Image resizing failed",
	},
	ErrContributionNew: {
		Message: "Post failed",
	},
	ErrContributionSave: {
		Message: "Failed to save",
	},
	ErrContributionTagSave: {
		Message: "Failed to save tag",
	},
	ErrUserSave: {
		Message: "Failed to save user",
	},
	ErrUserOrPasswordDifferent: {
		Message: "Email address and password do not match. Please enter again",
	},
	ErrContributionSearch: {
		Message: "Failed to get search results",
	},
	ErrFollowed: {
		Message: "Already followed",
	},
	ErrAddFollow: {
		Message: "Failed to register for following. Sorry to trouble you, but please add it again.",
	},
	ErrContributionNotFound: {
		Message: "Post data that does not exist",
	},
	ErrDeleteFollow: {
		Message: "Failed to delete follow. Please try again.",
	},
	ErrTagMaxNumberOver: {
		Message: "The number of tags that can be set has been exceeded. If you want to add one, please delete it.",
	},
	ErrTagNameOverlap: {
		Message: "The same tag is already registered",
	},
	ErrContributionNoUser: {
		Message: "You can not do this because it is not your own post",
	},
	ErrPasswordMinLength: {
		Message: "Please set a password with at least 8 characters.",
	},
}


// get ErroResponse get error response
func getErroResponse(errCode int) ErrorResponse {

	err := errResponseMap[errCode]
	err.ErrCode = errCode

	return err
}

// IsNoLogin Determine if you are logged in
func (c *BaseController) IsNoLogin(userID int) bool {
	if userID == noUserID {
		return false
	}

	return true
}

// ServerLoginNotFound I can not view without login
func (c *BaseController) ServerLoginNotFound() {
	c.ServerError (errors.New("login not found"), ErrCodeLoginNotFound, noUserID)
}

// ServerError Make server error
func (c *BaseController) ServerError(err error, errCode int, userID int) {
	beego.Error("Error:", err.Error())
	logs.Err(err.Error(), userID)

	c.Ctx.ResponseWriter.WriteHeader(500)
	c.Data["json"] = get ErroResponse(errCode)

	c.ServeJSON ()
}

// isTest test environment
func isTest() bool {
	if beego.AppConfig.String("runmode") == "test" {
		return true
	}

	return false
}

// RedirectError redirect error
func (c *BaseController) RedirectError(err error, userID int) {

	logs.Err(err.Error(), userID)

	c.Redirect(beego.AppConfig.String("errorUrl"), 302)
}
