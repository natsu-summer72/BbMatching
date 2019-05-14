// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// User service
//
// Command:
// $ goa gen github.com/natsu-summer72/BbMatching/design

package user

import (
	"context"

	userviews "github.com/natsu-summer72/BbMatching/gen/user/views"
	"goa.design/goa/security"
)

// ユーザー(チーム)に関するエンドポイント
type Service interface {
	// 現在のエンドポイントに紐づくユーザーの情報を返します。
	// The "view" return value must have one of the following views
	//	- "default"
	//	- "tiny"
	GetCurrentUser(context.Context, *SessionTokenPayload) (res *BbmatchingUser, view string, err error)
	// 指定したIDのユーザーの情報を取得します。
	// The "view" return value must have one of the following views
	//	- "default"
	//	- "tiny"
	GetUser(context.Context, *GetUserPayload) (res *BbmatchingUser, view string, err error)
	// ユーザーの一覧を取得します。
	// The "view" return value must have one of the following views
	//	- "default"
	//	- "tiny"
	ListUser(context.Context, *SessionTokenPayload) (res BbmatchingUserCollection, view string, err error)
	// 現在のセッションに紐づくユーザーの情報を更新します。
	// The "view" return value must have one of the following views
	//	- "default"
	//	- "tiny"
	UpdateCurrentUser(context.Context, *UpdateUserPayload) (res *BbmatchingUser, view string, err error)
	// 現在のセッションに紐づくユーザーを削除します。
	DeleteCurrentUser(context.Context, *SessionTokenPayload) (err error)
}

// Auther defines the authorization functions to be implemented by the service.
type Auther interface {
	// JWTAuth implements the authorization logic for the JWT security scheme.
	JWTAuth(ctx context.Context, token string, schema *security.JWTScheme) (context.Context, error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "User"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [5]string{"Get current user", "Get User", "List User", "Update current user", "Delete current user"}

// SessionTokenPayload is the payload type of the User service Get current user
// method.
type SessionTokenPayload struct {
	// JWT used for Authentication
	Token *string
}

// BbmatchingUser is the result type of the User service Get current user
// method.
type BbmatchingUser struct {
	// firebaseのユーザーID
	UserID string
	// チームのプライマリメールアドレス
	Email string
	// チームのメイン電話番号
	PhoneNumber string
	// チームの写真URL
	PhotoURL string
	// チームの表示名
	UserName string
}

// GetUserPayload is the payload type of the User service Get User method.
type GetUserPayload struct {
	// JWT used for Authentication
	Token *string
	// firebaseのユーザーID
	UserID string
}

// BbmatchingUserCollection is the result type of the User service List User
// method.
type BbmatchingUserCollection []*BbmatchingUser

// UpdateUserPayload is the payload type of the User service Update current
// user method.
type UpdateUserPayload struct {
	// JWT used for Authentication
	Token *string
	// チームのプライマリメールアドレス
	Email *string
	// チームのメイン電話番号
	PhoneNumber *string
	// チームの写真URL
	PhotoURL *string
	// チームの表示名
	UserName *string
}

// Credentials are invalid
type Unauthorized string

// Error returns an error description.
func (e Unauthorized) Error() string {
	return "Credentials are invalid"
}

// ErrorName returns "unauthorized".
func (e Unauthorized) ErrorName() string {
	return "unauthorized"
}

// NewBbmatchingUser initializes result type BbmatchingUser from viewed result
// type BbmatchingUser.
func NewBbmatchingUser(vres *userviews.BbmatchingUser) *BbmatchingUser {
	var res *BbmatchingUser
	switch vres.View {
	case "default", "":
		res = newBbmatchingUser(vres.Projected)
	case "tiny":
		res = newBbmatchingUserTiny(vres.Projected)
	}
	return res
}

// NewViewedBbmatchingUser initializes viewed result type BbmatchingUser from
// result type BbmatchingUser using the given view.
func NewViewedBbmatchingUser(res *BbmatchingUser, view string) *userviews.BbmatchingUser {
	var vres *userviews.BbmatchingUser
	switch view {
	case "default", "":
		p := newBbmatchingUserView(res)
		vres = &userviews.BbmatchingUser{p, "default"}
	case "tiny":
		p := newBbmatchingUserViewTiny(res)
		vres = &userviews.BbmatchingUser{p, "tiny"}
	}
	return vres
}

// NewBbmatchingUserCollection initializes result type BbmatchingUserCollection
// from viewed result type BbmatchingUserCollection.
func NewBbmatchingUserCollection(vres userviews.BbmatchingUserCollection) BbmatchingUserCollection {
	var res BbmatchingUserCollection
	switch vres.View {
	case "default", "":
		res = newBbmatchingUserCollection(vres.Projected)
	case "tiny":
		res = newBbmatchingUserCollectionTiny(vres.Projected)
	}
	return res
}

// NewViewedBbmatchingUserCollection initializes viewed result type
// BbmatchingUserCollection from result type BbmatchingUserCollection using the
// given view.
func NewViewedBbmatchingUserCollection(res BbmatchingUserCollection, view string) userviews.BbmatchingUserCollection {
	var vres userviews.BbmatchingUserCollection
	switch view {
	case "default", "":
		p := newBbmatchingUserCollectionView(res)
		vres = userviews.BbmatchingUserCollection{p, "default"}
	case "tiny":
		p := newBbmatchingUserCollectionViewTiny(res)
		vres = userviews.BbmatchingUserCollection{p, "tiny"}
	}
	return vres
}

// newBbmatchingUser converts projected type BbmatchingUser to service type
// BbmatchingUser.
func newBbmatchingUser(vres *userviews.BbmatchingUserView) *BbmatchingUser {
	res := &BbmatchingUser{}
	if vres.UserID != nil {
		res.UserID = *vres.UserID
	}
	if vres.Email != nil {
		res.Email = *vres.Email
	}
	if vres.PhoneNumber != nil {
		res.PhoneNumber = *vres.PhoneNumber
	}
	if vres.PhotoURL != nil {
		res.PhotoURL = *vres.PhotoURL
	}
	if vres.UserName != nil {
		res.UserName = *vres.UserName
	}
	return res
}

// newBbmatchingUserTiny converts projected type BbmatchingUser to service type
// BbmatchingUser.
func newBbmatchingUserTiny(vres *userviews.BbmatchingUserView) *BbmatchingUser {
	res := &BbmatchingUser{}
	if vres.UserName != nil {
		res.UserName = *vres.UserName
	}
	if vres.Email != nil {
		res.Email = *vres.Email
	}
	if vres.PhotoURL != nil {
		res.PhotoURL = *vres.PhotoURL
	}
	return res
}

// newBbmatchingUserView projects result type BbmatchingUser into projected
// type BbmatchingUserView using the "default" view.
func newBbmatchingUserView(res *BbmatchingUser) *userviews.BbmatchingUserView {
	vres := &userviews.BbmatchingUserView{
		UserID:      &res.UserID,
		Email:       &res.Email,
		PhoneNumber: &res.PhoneNumber,
		PhotoURL:    &res.PhotoURL,
		UserName:    &res.UserName,
	}
	return vres
}

// newBbmatchingUserViewTiny projects result type BbmatchingUser into projected
// type BbmatchingUserView using the "tiny" view.
func newBbmatchingUserViewTiny(res *BbmatchingUser) *userviews.BbmatchingUserView {
	vres := &userviews.BbmatchingUserView{
		Email:    &res.Email,
		PhotoURL: &res.PhotoURL,
		UserName: &res.UserName,
	}
	return vres
}

// newBbmatchingUserCollection converts projected type BbmatchingUserCollection
// to service type BbmatchingUserCollection.
func newBbmatchingUserCollection(vres userviews.BbmatchingUserCollectionView) BbmatchingUserCollection {
	res := make(BbmatchingUserCollection, len(vres))
	for i, n := range vres {
		res[i] = newBbmatchingUser(n)
	}
	return res
}

// newBbmatchingUserCollectionTiny converts projected type
// BbmatchingUserCollection to service type BbmatchingUserCollection.
func newBbmatchingUserCollectionTiny(vres userviews.BbmatchingUserCollectionView) BbmatchingUserCollection {
	res := make(BbmatchingUserCollection, len(vres))
	for i, n := range vres {
		res[i] = newBbmatchingUserTiny(n)
	}
	return res
}

// newBbmatchingUserCollectionView projects result type
// BbmatchingUserCollection into projected type BbmatchingUserCollectionView
// using the "default" view.
func newBbmatchingUserCollectionView(res BbmatchingUserCollection) userviews.BbmatchingUserCollectionView {
	vres := make(userviews.BbmatchingUserCollectionView, len(res))
	for i, n := range res {
		vres[i] = newBbmatchingUserView(n)
	}
	return vres
}

// newBbmatchingUserCollectionViewTiny projects result type
// BbmatchingUserCollection into projected type BbmatchingUserCollectionView
// using the "tiny" view.
func newBbmatchingUserCollectionViewTiny(res BbmatchingUserCollection) userviews.BbmatchingUserCollectionView {
	vres := make(userviews.BbmatchingUserCollectionView, len(res))
	for i, n := range res {
		vres[i] = newBbmatchingUserViewTiny(n)
	}
	return vres
}