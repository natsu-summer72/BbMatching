// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// User HTTP client encoders and decoders
//
// Command:
// $ goa gen github.com/natsu-summer72/BbMatching/design

package client

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"

	user "github.com/natsu-summer72/BbMatching/gen/user"
	userviews "github.com/natsu-summer72/BbMatching/gen/user/views"
	goahttp "goa.design/goa/http"
)

// BuildGetCurrentUserRequest instantiates a HTTP request object with method
// and path set to call the "User" service "Get current user" endpoint
func (c *Client) BuildGetCurrentUserRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: GetCurrentUserUserPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("User", "Get current user", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeGetCurrentUserRequest returns an encoder for requests sent to the User
// Get current user server.
func EncodeGetCurrentUserRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*user.SessionTokenPayload)
		if !ok {
			return goahttp.ErrInvalidType("User", "Get current user", "*user.SessionTokenPayload", v)
		}
		if p.Token != nil {
			req.Header.Set("Authorization", *p.Token)
		}
		return nil
	}
}

// DecodeGetCurrentUserResponse returns a decoder for responses returned by the
// User Get current user endpoint. restoreBody controls whether the response
// body should be restored after having been read.
// DecodeGetCurrentUserResponse may return the following errors:
//	- "unauthorized" (type user.Unauthorized): http.StatusUnauthorized
//	- error: internal error
func DecodeGetCurrentUserResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body GetCurrentUserResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("User", "Get current user", err)
			}
			p := NewGetCurrentUserBbmatchingUserOK(&body)
			view := resp.Header.Get("goa-view")
			vres := &userviews.BbmatchingUser{p, view}
			if err = userviews.ValidateBbmatchingUser(vres); err != nil {
				return nil, goahttp.ErrValidationError("User", "Get current user", err)
			}
			res := user.NewBbmatchingUser(vres)
			return res, nil
		case http.StatusUnauthorized:
			var (
				body GetCurrentUserUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("User", "Get current user", err)
			}
			return nil, NewGetCurrentUserUnauthorized(body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("User", "Get current user", resp.StatusCode, string(body))
		}
	}
}

// BuildGetUserRequest instantiates a HTTP request object with method and path
// set to call the "User" service "Get User" endpoint
func (c *Client) BuildGetUserRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		userID string
	)
	{
		p, ok := v.(*user.GetUserPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("User", "Get User", "*user.GetUserPayload", v)
		}
		userID = p.UserID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: GetUserUserPath(userID)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("User", "Get User", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeGetUserRequest returns an encoder for requests sent to the User Get
// User server.
func EncodeGetUserRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*user.GetUserPayload)
		if !ok {
			return goahttp.ErrInvalidType("User", "Get User", "*user.GetUserPayload", v)
		}
		if p.Token != nil {
			req.Header.Set("Authorization", *p.Token)
		}
		return nil
	}
}

// DecodeGetUserResponse returns a decoder for responses returned by the User
// Get User endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeGetUserResponse may return the following errors:
//	- "unauthorized" (type user.Unauthorized): http.StatusUnauthorized
//	- error: internal error
func DecodeGetUserResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body GetUserResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("User", "Get User", err)
			}
			p := NewGetUserBbmatchingUserOK(&body)
			view := resp.Header.Get("goa-view")
			vres := &userviews.BbmatchingUser{p, view}
			if err = userviews.ValidateBbmatchingUser(vres); err != nil {
				return nil, goahttp.ErrValidationError("User", "Get User", err)
			}
			res := user.NewBbmatchingUser(vres)
			return res, nil
		case http.StatusUnauthorized:
			var (
				body GetUserUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("User", "Get User", err)
			}
			return nil, NewGetUserUnauthorized(body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("User", "Get User", resp.StatusCode, string(body))
		}
	}
}

// BuildListUserRequest instantiates a HTTP request object with method and path
// set to call the "User" service "List User" endpoint
func (c *Client) BuildListUserRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ListUserUserPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("User", "List User", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeListUserRequest returns an encoder for requests sent to the User List
// User server.
func EncodeListUserRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*user.SessionTokenPayload)
		if !ok {
			return goahttp.ErrInvalidType("User", "List User", "*user.SessionTokenPayload", v)
		}
		if p.Token != nil {
			req.Header.Set("Authorization", *p.Token)
		}
		return nil
	}
}

// DecodeListUserResponse returns a decoder for responses returned by the User
// List User endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeListUserResponse may return the following errors:
//	- "unauthorized" (type user.Unauthorized): http.StatusUnauthorized
//	- error: internal error
func DecodeListUserResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body ListUserResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("User", "List User", err)
			}
			p := NewListUserBbmatchingUserCollectionOK(body)
			view := resp.Header.Get("goa-view")
			vres := userviews.BbmatchingUserCollection{p, view}
			if err = userviews.ValidateBbmatchingUserCollection(vres); err != nil {
				return nil, goahttp.ErrValidationError("User", "List User", err)
			}
			res := user.NewBbmatchingUserCollection(vres)
			return res, nil
		case http.StatusUnauthorized:
			var (
				body ListUserUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("User", "List User", err)
			}
			return nil, NewListUserUnauthorized(body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("User", "List User", resp.StatusCode, string(body))
		}
	}
}

// BuildUpdateCurrentUserRequest instantiates a HTTP request object with method
// and path set to call the "User" service "Update current user" endpoint
func (c *Client) BuildUpdateCurrentUserRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: UpdateCurrentUserUserPath()}
	req, err := http.NewRequest("PUT", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("User", "Update current user", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeUpdateCurrentUserRequest returns an encoder for requests sent to the
// User Update current user server.
func EncodeUpdateCurrentUserRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*user.UpdateUserPayload)
		if !ok {
			return goahttp.ErrInvalidType("User", "Update current user", "*user.UpdateUserPayload", v)
		}
		if p.Token != nil {
			req.Header.Set("Authorization", *p.Token)
		}
		body := NewUpdateCurrentUserRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("User", "Update current user", err)
		}
		return nil
	}
}

// DecodeUpdateCurrentUserResponse returns a decoder for responses returned by
// the User Update current user endpoint. restoreBody controls whether the
// response body should be restored after having been read.
// DecodeUpdateCurrentUserResponse may return the following errors:
//	- "unauthorized" (type user.Unauthorized): http.StatusUnauthorized
//	- error: internal error
func DecodeUpdateCurrentUserResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body UpdateCurrentUserResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("User", "Update current user", err)
			}
			p := NewUpdateCurrentUserBbmatchingUserOK(&body)
			view := resp.Header.Get("goa-view")
			vres := &userviews.BbmatchingUser{p, view}
			if err = userviews.ValidateBbmatchingUser(vres); err != nil {
				return nil, goahttp.ErrValidationError("User", "Update current user", err)
			}
			res := user.NewBbmatchingUser(vres)
			return res, nil
		case http.StatusUnauthorized:
			var (
				body UpdateCurrentUserUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("User", "Update current user", err)
			}
			return nil, NewUpdateCurrentUserUnauthorized(body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("User", "Update current user", resp.StatusCode, string(body))
		}
	}
}

// BuildDeleteCurrentUserRequest instantiates a HTTP request object with method
// and path set to call the "User" service "Delete current user" endpoint
func (c *Client) BuildDeleteCurrentUserRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: DeleteCurrentUserUserPath()}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("User", "Delete current user", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeDeleteCurrentUserRequest returns an encoder for requests sent to the
// User Delete current user server.
func EncodeDeleteCurrentUserRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*user.SessionTokenPayload)
		if !ok {
			return goahttp.ErrInvalidType("User", "Delete current user", "*user.SessionTokenPayload", v)
		}
		if p.Token != nil {
			req.Header.Set("Authorization", *p.Token)
		}
		return nil
	}
}

// DecodeDeleteCurrentUserResponse returns a decoder for responses returned by
// the User Delete current user endpoint. restoreBody controls whether the
// response body should be restored after having been read.
// DecodeDeleteCurrentUserResponse may return the following errors:
//	- "unauthorized" (type user.Unauthorized): http.StatusUnauthorized
//	- error: internal error
func DecodeDeleteCurrentUserResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusNoContent:
			return nil, nil
		case http.StatusUnauthorized:
			var (
				body DeleteCurrentUserUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("User", "Delete current user", err)
			}
			return nil, NewDeleteCurrentUserUnauthorized(body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("User", "Delete current user", resp.StatusCode, string(body))
		}
	}
}
