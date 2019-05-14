package bbm

import (
	user "github.com/natsu-summer72/BbMatching/gen/user"
	"context"
	"log"

	"firebase.google.com/go/auth"
)

// User service example implementation.
// The example methods log the requests and return zero values.
type usersrvc struct {
	logger *log.Logger
	authClient *auth.Client
}

// NewUser returns the User service implementation.
func NewUser(logger *log.Logger, client *auth.Client) user.Service {
	return &usersrvc{logger, client}
}

// 現在のエンドポイントに紐づくユーザーの情報を返します。
func (s *usersrvc) GetCurrentUser(ctx context.Context, p *user.SessionTokenPayload) (res *user.BbmatchingUser, view string, err error) {
	res = &user.BbmatchingUser{}
	view = "default"
	s.logger.Print("user.Get current user")

	verifiedToken, err := s.authClient.VerifyIDToken(ctx, *p.Token)
	u, err := s.authClient.GetUser(ctx, verifiedToken.UID)
	if err != nil {
		log.Fatalf("error getting Auth Client: %v\n", err)
		return
	}
	res = &user.BbmatchingUser{
		UserID: u.UID,
		Email: u.Email,
		PhoneNumber:u.PhoneNumber,
		PhotoURL:u.PhotoURL,
		UserName:u.DisplayName,
	}

	return
}

// 指定したIDのユーザーの情報を取得します。
func (s *usersrvc) GetUser(ctx context.Context, p *user.GetUserPayload) (res *user.BbmatchingUser, view string, err error) {
	res = &user.BbmatchingUser{}
	view = "default"
	s.logger.Print("user.Get User")
	return
}

// ユーザーの一覧を取得します。
func (s *usersrvc) ListUser(ctx context.Context, p *user.SessionTokenPayload) (res user.BbmatchingUserCollection, view string, err error) {
	view = "default"
	s.logger.Print("user.List User")
	return
}

// 現在のセッションに紐づくユーザーの情報を更新します。
func (s *usersrvc) UpdateCurrentUser(ctx context.Context, p *user.UpdateUserPayload) (res *user.BbmatchingUser, view string, err error) {
	res = &user.BbmatchingUser{}
	view = "default"
	s.logger.Print("user.Update current user")
	return
}

// 現在のセッションに紐づくユーザーを削除します。
func (s *usersrvc) DeleteCurrentUser(ctx context.Context, p *user.SessionTokenPayload) (err error) {
	s.logger.Print("user.Delete current user")
	return
}
