package bbm

import (
	"bytes"
	"context"
	"encoding/json"
	"firebase.google.com/go/auth"
	"fmt"
	"github.com/natsu-summer72/BbMatching/gen/user"
	"google.golang.org/api/iterator"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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
		EmailVerified:u.EmailVerified,
	}

	return
}

// 指定したIDのユーザーの情報を取得します。
func (s *usersrvc) GetUser(ctx context.Context, p *user.GetUserPayload) (res *user.BbmatchingUser, view string, err error) {
	res = &user.BbmatchingUser{}
	view = "default"
	s.logger.Print("user.Get User")

	u, err := s.authClient.GetUser(ctx, p.UserID)
	if err!=nil{
		return
	}
	res = &user.BbmatchingUser{
		UserID: u.UID,
		UserName:u.DisplayName,
		Email:u.Email,
		PhoneNumber: u.PhoneNumber,
		PhotoURL: u.PhotoURL,
		EmailVerified: u.EmailVerified,
	}

	return
}

// ユーザーの一覧を取得します。
func (s *usersrvc) ListUser(ctx context.Context, p *user.SessionTokenPayload) (res user.BbmatchingUserCollection, view string, err error) {
	view = "default"
	s.logger.Print("user.List User")

	iter := s.authClient.Users(ctx, "")
	for {
		u, err := iter.Next()
		if err != nil {
			if err == iterator.Done {
				break
			}
			return nil, view, err
		}
		r := &user.BbmatchingUser{
			UserID: u.UID,
			UserName: u.DisplayName,
			Email: u.Email,
			PhoneNumber: u.PhoneNumber,
			PhotoURL: u.PhotoURL,
			EmailVerified: u.EmailVerified,
		}
		res = append(res, r)
	}
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

// 指定したユーザーIDのJWTを取得します
func (s *usersrvc) GetJWT(ctx context.Context, p *user.GetJWTPayload) (res *user.BbmatchingJWT, err error) {
	res = &user.BbmatchingJWT{}
	s.logger.Print("user.Get JWT")
	token, err := s.authClient.CustomToken(ctx, p.UserID)

	if err!=nil{
		log.Fatalf("error minting custom token")
	}
	IDtoken, err := signInWithCustomToken(token)

	res = &user.BbmatchingJWT{
		JWT: &IDtoken,
	}

	return
}


// firebaseのカスタムトークンから，IDTokenを取得
func signInWithCustomToken(token string) (string, error) {
	req, err := json.Marshal(map[string]interface{}{
		"token":             token,
		"returnSecureToken": true,
	})
	if err != nil {
		return "", err
	}

	path := fmt.Sprintf("https://www.googleapis.com/identitytoolkit/v3/relyingparty/verifyCustomToken?key=%s", os.Getenv("FIREBASE_APIKEY") )
	res, err := http.Post(path, "application/json", bytes.NewBuffer(req))
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected http status code: %d", res.StatusCode)
	}
	resp, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	var resBody struct {
		IDToken string `json:"idToken"`
	}
	if err := json.Unmarshal(resp, &resBody); err != nil {
		return "", err
	}
	return resBody.IDToken, err
}
