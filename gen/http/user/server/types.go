// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// User HTTP server types
//
// Command:
// $ goa gen github.com/natsu-summer72/BbMatching/design

package server

import (
	"unicode/utf8"

	user "github.com/natsu-summer72/BbMatching/gen/user"
	userviews "github.com/natsu-summer72/BbMatching/gen/user/views"
	goa "goa.design/goa"
)

// UpdateCurrentUserRequestBody is the type of the "User" service "Update
// current user" endpoint HTTP request body.
type UpdateCurrentUserRequestBody struct {
	// firebaseのユーザーID
	UserID *string `form:"user_id,omitempty" json:"user_id,omitempty" xml:"user_id,omitempty"`
	// チームのプライマリメールアドレス
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// チームのメイン電話番号
	PhoneNumber *string `form:"phoneNumber,omitempty" json:"phoneNumber,omitempty" xml:"phoneNumber,omitempty"`
	// チームの写真URL
	PhotoURL *string `form:"photoURL,omitempty" json:"photoURL,omitempty" xml:"photoURL,omitempty"`
	// チームの表示名
	UserName *string `form:"UserName,omitempty" json:"UserName,omitempty" xml:"UserName,omitempty"`
}

// GetCurrentUserResponseBody is the type of the "User" service "Get current
// user" endpoint HTTP response body.
type GetCurrentUserResponseBody struct {
	// firebaseのユーザーID
	UserID string `form:"user_id" json:"user_id" xml:"user_id"`
	// チームのプライマリメールアドレス
	Email string `form:"email" json:"email" xml:"email"`
	// チームのメイン電話番号
	PhoneNumber string `form:"phoneNumber" json:"phoneNumber" xml:"phoneNumber"`
	// チームの写真URL
	PhotoURL string `form:"photoURL" json:"photoURL" xml:"photoURL"`
	// チームの表示名
	UserName string `form:"UserName" json:"UserName" xml:"UserName"`
	// ユーザーのプライマリメールアドレスが確認されているか。
	EmailVerified bool `form:"email_verified" json:"email_verified" xml:"email_verified"`
}

// GetCurrentUserResponseBodyTiny is the type of the "User" service "Get
// current user" endpoint HTTP response body.
type GetCurrentUserResponseBodyTiny struct {
	// チームの表示名
	UserName string `form:"UserName" json:"UserName" xml:"UserName"`
	// チームのプライマリメールアドレス
	Email string `form:"email" json:"email" xml:"email"`
	// チームの写真URL
	PhotoURL string `form:"photoURL" json:"photoURL" xml:"photoURL"`
}

// GetUserResponseBody is the type of the "User" service "Get User" endpoint
// HTTP response body.
type GetUserResponseBody struct {
	// firebaseのユーザーID
	UserID string `form:"user_id" json:"user_id" xml:"user_id"`
	// チームのプライマリメールアドレス
	Email string `form:"email" json:"email" xml:"email"`
	// チームのメイン電話番号
	PhoneNumber string `form:"phoneNumber" json:"phoneNumber" xml:"phoneNumber"`
	// チームの写真URL
	PhotoURL string `form:"photoURL" json:"photoURL" xml:"photoURL"`
	// チームの表示名
	UserName string `form:"UserName" json:"UserName" xml:"UserName"`
	// ユーザーのプライマリメールアドレスが確認されているか。
	EmailVerified bool `form:"email_verified" json:"email_verified" xml:"email_verified"`
}

// GetUserResponseBodyTiny is the type of the "User" service "Get User"
// endpoint HTTP response body.
type GetUserResponseBodyTiny struct {
	// チームの表示名
	UserName string `form:"UserName" json:"UserName" xml:"UserName"`
	// チームのプライマリメールアドレス
	Email string `form:"email" json:"email" xml:"email"`
	// チームの写真URL
	PhotoURL string `form:"photoURL" json:"photoURL" xml:"photoURL"`
}

// BbmatchingUserResponseCollection is the type of the "User" service "List
// User" endpoint HTTP response body.
type BbmatchingUserResponseCollection []*BbmatchingUserResponse

// BbmatchingUserResponseTinyCollection is the type of the "User" service "List
// User" endpoint HTTP response body.
type BbmatchingUserResponseTinyCollection []*BbmatchingUserResponseTiny

// UpdateCurrentUserResponseBody is the type of the "User" service "Update
// current user" endpoint HTTP response body.
type UpdateCurrentUserResponseBody struct {
	// firebaseのユーザーID
	UserID string `form:"user_id" json:"user_id" xml:"user_id"`
	// チームのプライマリメールアドレス
	Email string `form:"email" json:"email" xml:"email"`
	// チームのメイン電話番号
	PhoneNumber string `form:"phoneNumber" json:"phoneNumber" xml:"phoneNumber"`
	// チームの写真URL
	PhotoURL string `form:"photoURL" json:"photoURL" xml:"photoURL"`
	// チームの表示名
	UserName string `form:"UserName" json:"UserName" xml:"UserName"`
	// ユーザーのプライマリメールアドレスが確認されているか。
	EmailVerified bool `form:"email_verified" json:"email_verified" xml:"email_verified"`
}

// UpdateCurrentUserResponseBodyTiny is the type of the "User" service "Update
// current user" endpoint HTTP response body.
type UpdateCurrentUserResponseBodyTiny struct {
	// チームの表示名
	UserName string `form:"UserName" json:"UserName" xml:"UserName"`
	// チームのプライマリメールアドレス
	Email string `form:"email" json:"email" xml:"email"`
	// チームの写真URL
	PhotoURL string `form:"photoURL" json:"photoURL" xml:"photoURL"`
}

// GetJWTResponseBody is the type of the "User" service "Get JWT" endpoint HTTP
// response body.
type GetJWTResponseBody struct {
	// Json Web Token
	JWT *string `form:"JWT,omitempty" json:"JWT,omitempty" xml:"JWT,omitempty"`
}

// GetCurrentUserUnauthorizedResponseBody is the type of the "User" service
// "Get current user" endpoint HTTP response body for the "unauthorized" error.
type GetCurrentUserUnauthorizedResponseBody string

// GetUserUnauthorizedResponseBody is the type of the "User" service "Get User"
// endpoint HTTP response body for the "unauthorized" error.
type GetUserUnauthorizedResponseBody string

// ListUserUnauthorizedResponseBody is the type of the "User" service "List
// User" endpoint HTTP response body for the "unauthorized" error.
type ListUserUnauthorizedResponseBody string

// UpdateCurrentUserUnauthorizedResponseBody is the type of the "User" service
// "Update current user" endpoint HTTP response body for the "unauthorized"
// error.
type UpdateCurrentUserUnauthorizedResponseBody string

// DeleteCurrentUserUnauthorizedResponseBody is the type of the "User" service
// "Delete current user" endpoint HTTP response body for the "unauthorized"
// error.
type DeleteCurrentUserUnauthorizedResponseBody string

// GetJWTUnauthorizedResponseBody is the type of the "User" service "Get JWT"
// endpoint HTTP response body for the "unauthorized" error.
type GetJWTUnauthorizedResponseBody string

// BbmatchingUserResponse is used to define fields on response body types.
type BbmatchingUserResponse struct {
	// firebaseのユーザーID
	UserID string `form:"user_id" json:"user_id" xml:"user_id"`
	// チームのプライマリメールアドレス
	Email string `form:"email" json:"email" xml:"email"`
	// チームのメイン電話番号
	PhoneNumber string `form:"phoneNumber" json:"phoneNumber" xml:"phoneNumber"`
	// チームの写真URL
	PhotoURL string `form:"photoURL" json:"photoURL" xml:"photoURL"`
	// チームの表示名
	UserName string `form:"UserName" json:"UserName" xml:"UserName"`
	// ユーザーのプライマリメールアドレスが確認されているか。
	EmailVerified bool `form:"email_verified" json:"email_verified" xml:"email_verified"`
}

// BbmatchingUserResponseTiny is used to define fields on response body types.
type BbmatchingUserResponseTiny struct {
	// チームの表示名
	UserName string `form:"UserName" json:"UserName" xml:"UserName"`
	// チームのプライマリメールアドレス
	Email string `form:"email" json:"email" xml:"email"`
	// チームの写真URL
	PhotoURL string `form:"photoURL" json:"photoURL" xml:"photoURL"`
}

// NewGetCurrentUserResponseBody builds the HTTP response body from the result
// of the "Get current user" endpoint of the "User" service.
func NewGetCurrentUserResponseBody(res *userviews.BbmatchingUserView) *GetCurrentUserResponseBody {
	body := &GetCurrentUserResponseBody{
		UserID:        *res.UserID,
		UserName:      *res.UserName,
		Email:         *res.Email,
		PhoneNumber:   *res.PhoneNumber,
		PhotoURL:      *res.PhotoURL,
		EmailVerified: *res.EmailVerified,
	}
	return body
}

// NewGetCurrentUserResponseBodyTiny builds the HTTP response body from the
// result of the "Get current user" endpoint of the "User" service.
func NewGetCurrentUserResponseBodyTiny(res *userviews.BbmatchingUserView) *GetCurrentUserResponseBodyTiny {
	body := &GetCurrentUserResponseBodyTiny{
		UserName: *res.UserName,
		Email:    *res.Email,
		PhotoURL: *res.PhotoURL,
	}
	return body
}

// NewGetUserResponseBody builds the HTTP response body from the result of the
// "Get User" endpoint of the "User" service.
func NewGetUserResponseBody(res *userviews.BbmatchingUserView) *GetUserResponseBody {
	body := &GetUserResponseBody{
		UserID:        *res.UserID,
		UserName:      *res.UserName,
		Email:         *res.Email,
		PhoneNumber:   *res.PhoneNumber,
		PhotoURL:      *res.PhotoURL,
		EmailVerified: *res.EmailVerified,
	}
	return body
}

// NewGetUserResponseBodyTiny builds the HTTP response body from the result of
// the "Get User" endpoint of the "User" service.
func NewGetUserResponseBodyTiny(res *userviews.BbmatchingUserView) *GetUserResponseBodyTiny {
	body := &GetUserResponseBodyTiny{
		UserName: *res.UserName,
		Email:    *res.Email,
		PhotoURL: *res.PhotoURL,
	}
	return body
}

// NewBbmatchingUserResponseCollection builds the HTTP response body from the
// result of the "List User" endpoint of the "User" service.
func NewBbmatchingUserResponseCollection(res userviews.BbmatchingUserCollectionView) BbmatchingUserResponseCollection {
	body := make([]*BbmatchingUserResponse, len(res))
	for i, val := range res {
		body[i] = &BbmatchingUserResponse{
			UserID:        *val.UserID,
			UserName:      *val.UserName,
			Email:         *val.Email,
			PhoneNumber:   *val.PhoneNumber,
			PhotoURL:      *val.PhotoURL,
			EmailVerified: *val.EmailVerified,
		}
	}
	return body
}

// NewBbmatchingUserResponseTinyCollection builds the HTTP response body from
// the result of the "List User" endpoint of the "User" service.
func NewBbmatchingUserResponseTinyCollection(res userviews.BbmatchingUserCollectionView) BbmatchingUserResponseTinyCollection {
	body := make([]*BbmatchingUserResponseTiny, len(res))
	for i, val := range res {
		body[i] = &BbmatchingUserResponseTiny{
			UserName: *val.UserName,
			Email:    *val.Email,
			PhotoURL: *val.PhotoURL,
		}
	}
	return body
}

// NewUpdateCurrentUserResponseBody builds the HTTP response body from the
// result of the "Update current user" endpoint of the "User" service.
func NewUpdateCurrentUserResponseBody(res *userviews.BbmatchingUserView) *UpdateCurrentUserResponseBody {
	body := &UpdateCurrentUserResponseBody{
		UserID:        *res.UserID,
		UserName:      *res.UserName,
		Email:         *res.Email,
		PhoneNumber:   *res.PhoneNumber,
		PhotoURL:      *res.PhotoURL,
		EmailVerified: *res.EmailVerified,
	}
	return body
}

// NewUpdateCurrentUserResponseBodyTiny builds the HTTP response body from the
// result of the "Update current user" endpoint of the "User" service.
func NewUpdateCurrentUserResponseBodyTiny(res *userviews.BbmatchingUserView) *UpdateCurrentUserResponseBodyTiny {
	body := &UpdateCurrentUserResponseBodyTiny{
		UserName: *res.UserName,
		Email:    *res.Email,
		PhotoURL: *res.PhotoURL,
	}
	return body
}

// NewGetJWTResponseBody builds the HTTP response body from the result of the
// "Get JWT" endpoint of the "User" service.
func NewGetJWTResponseBody(res *userviews.BbmatchingJWTView) *GetJWTResponseBody {
	body := &GetJWTResponseBody{
		JWT: res.JWT,
	}
	return body
}

// NewGetCurrentUserUnauthorizedResponseBody builds the HTTP response body from
// the result of the "Get current user" endpoint of the "User" service.
func NewGetCurrentUserUnauthorizedResponseBody(res user.Unauthorized) GetCurrentUserUnauthorizedResponseBody {
	body := GetCurrentUserUnauthorizedResponseBody(res)
	return body
}

// NewGetUserUnauthorizedResponseBody builds the HTTP response body from the
// result of the "Get User" endpoint of the "User" service.
func NewGetUserUnauthorizedResponseBody(res user.Unauthorized) GetUserUnauthorizedResponseBody {
	body := GetUserUnauthorizedResponseBody(res)
	return body
}

// NewListUserUnauthorizedResponseBody builds the HTTP response body from the
// result of the "List User" endpoint of the "User" service.
func NewListUserUnauthorizedResponseBody(res user.Unauthorized) ListUserUnauthorizedResponseBody {
	body := ListUserUnauthorizedResponseBody(res)
	return body
}

// NewUpdateCurrentUserUnauthorizedResponseBody builds the HTTP response body
// from the result of the "Update current user" endpoint of the "User" service.
func NewUpdateCurrentUserUnauthorizedResponseBody(res user.Unauthorized) UpdateCurrentUserUnauthorizedResponseBody {
	body := UpdateCurrentUserUnauthorizedResponseBody(res)
	return body
}

// NewDeleteCurrentUserUnauthorizedResponseBody builds the HTTP response body
// from the result of the "Delete current user" endpoint of the "User" service.
func NewDeleteCurrentUserUnauthorizedResponseBody(res user.Unauthorized) DeleteCurrentUserUnauthorizedResponseBody {
	body := DeleteCurrentUserUnauthorizedResponseBody(res)
	return body
}

// NewGetJWTUnauthorizedResponseBody builds the HTTP response body from the
// result of the "Get JWT" endpoint of the "User" service.
func NewGetJWTUnauthorizedResponseBody(res user.Unauthorized) GetJWTUnauthorizedResponseBody {
	body := GetJWTUnauthorizedResponseBody(res)
	return body
}

// NewGetCurrentUserSessionTokenPayload builds a User service Get current user
// endpoint payload.
func NewGetCurrentUserSessionTokenPayload(token *string) *user.SessionTokenPayload {
	return &user.SessionTokenPayload{
		Token: token,
	}
}

// NewGetUserPayload builds a User service Get User endpoint payload.
func NewGetUserPayload(userID string, token *string) *user.GetUserPayload {
	return &user.GetUserPayload{
		UserID: userID,
		Token:  token,
	}
}

// NewListUserSessionTokenPayload builds a User service List User endpoint
// payload.
func NewListUserSessionTokenPayload(token *string) *user.SessionTokenPayload {
	return &user.SessionTokenPayload{
		Token: token,
	}
}

// NewUpdateCurrentUserUpdateUserPayload builds a User service Update current
// user endpoint payload.
func NewUpdateCurrentUserUpdateUserPayload(body *UpdateCurrentUserRequestBody, token *string) *user.UpdateUserPayload {
	v := &user.UpdateUserPayload{
		UserID:      body.UserID,
		Email:       body.Email,
		PhoneNumber: body.PhoneNumber,
		PhotoURL:    body.PhotoURL,
		UserName:    body.UserName,
	}
	v.Token = token
	return v
}

// NewDeleteCurrentUserSessionTokenPayload builds a User service Delete current
// user endpoint payload.
func NewDeleteCurrentUserSessionTokenPayload(token *string) *user.SessionTokenPayload {
	return &user.SessionTokenPayload{
		Token: token,
	}
}

// NewGetJWTPayload builds a User service Get JWT endpoint payload.
func NewGetJWTPayload(userID string) *user.GetJWTPayload {
	return &user.GetJWTPayload{
		UserID: userID,
	}
}

// ValidateUpdateCurrentUserRequestBody runs the validations defined on Update
// Current UserRequestBody
func ValidateUpdateCurrentUserRequestBody(body *UpdateCurrentUserRequestBody) (err error) {
	if body.UserID != nil {
		if utf8.RuneCountInString(*body.UserID) < 28 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.user_id", *body.UserID, utf8.RuneCountInString(*body.UserID), 28, true))
		}
	}
	if body.UserID != nil {
		if utf8.RuneCountInString(*body.UserID) > 28 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.user_id", *body.UserID, utf8.RuneCountInString(*body.UserID), 28, false))
		}
	}
	if body.Email != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.email", *body.Email, goa.FormatEmail))
	}
	if body.PhoneNumber != nil {
		err = goa.MergeErrors(err, goa.ValidatePattern("body.phoneNumber", *body.PhoneNumber, "^\\+?[\\d]{10,}$"))
	}
	return
}

// ValidateBbmatchingUserResponse runs the validations defined on
// BbmatchingUserResponse
func ValidateBbmatchingUserResponse(body *BbmatchingUserResponse) (err error) {
	if utf8.RuneCountInString(body.UserID) < 28 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("body.user_id", body.UserID, utf8.RuneCountInString(body.UserID), 28, true))
	}
	if utf8.RuneCountInString(body.UserID) > 28 {
		err = goa.MergeErrors(err, goa.InvalidLengthError("body.user_id", body.UserID, utf8.RuneCountInString(body.UserID), 28, false))
	}
	err = goa.MergeErrors(err, goa.ValidateFormat("body.email", body.Email, goa.FormatEmail))

	err = goa.MergeErrors(err, goa.ValidatePattern("body.phoneNumber", body.PhoneNumber, "^\\+?[\\d]{10,}$"))
	return
}

// ValidateBbmatchingUserResponseTiny runs the validations defined on
// BbmatchingUserResponseTiny
func ValidateBbmatchingUserResponseTiny(body *BbmatchingUserResponseTiny) (err error) {
	err = goa.MergeErrors(err, goa.ValidateFormat("body.email", body.Email, goa.FormatEmail))

	return
}
