// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// User client HTTP transport
//
// Command:
// $ goa gen github.com/natsu-summer72/BbMatching/design

package client

import (
	"context"
	"net/http"

	goa "goa.design/goa"
	goahttp "goa.design/goa/http"
)

// Client lists the User service endpoint HTTP clients.
type Client struct {
	// GetCurrentUser Doer is the HTTP client used to make requests to the Get
	// current user endpoint.
	GetCurrentUserDoer goahttp.Doer

	// GetUser Doer is the HTTP client used to make requests to the Get User
	// endpoint.
	GetUserDoer goahttp.Doer

	// ListUser Doer is the HTTP client used to make requests to the List User
	// endpoint.
	ListUserDoer goahttp.Doer

	// UpdateCurrentUser Doer is the HTTP client used to make requests to the
	// Update current user endpoint.
	UpdateCurrentUserDoer goahttp.Doer

	// DeleteCurrentUser Doer is the HTTP client used to make requests to the
	// Delete current user endpoint.
	DeleteCurrentUserDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the User service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		GetCurrentUserDoer:    doer,
		GetUserDoer:           doer,
		ListUserDoer:          doer,
		UpdateCurrentUserDoer: doer,
		DeleteCurrentUserDoer: doer,
		RestoreResponseBody:   restoreBody,
		scheme:                scheme,
		host:                  host,
		decoder:               dec,
		encoder:               enc,
	}
}

// GetCurrentUser returns an endpoint that makes HTTP requests to the User
// service Get current user server.
func (c *Client) GetCurrentUser() goa.Endpoint {
	var (
		encodeRequest  = EncodeGetCurrentUserRequest(c.encoder)
		decodeResponse = DecodeGetCurrentUserResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildGetCurrentUserRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.GetCurrentUserDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("User", "Get current user", err)
		}
		return decodeResponse(resp)
	}
}

// GetUser returns an endpoint that makes HTTP requests to the User service Get
// User server.
func (c *Client) GetUser() goa.Endpoint {
	var (
		encodeRequest  = EncodeGetUserRequest(c.encoder)
		decodeResponse = DecodeGetUserResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildGetUserRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.GetUserDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("User", "Get User", err)
		}
		return decodeResponse(resp)
	}
}

// ListUser returns an endpoint that makes HTTP requests to the User service
// List User server.
func (c *Client) ListUser() goa.Endpoint {
	var (
		encodeRequest  = EncodeListUserRequest(c.encoder)
		decodeResponse = DecodeListUserResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildListUserRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ListUserDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("User", "List User", err)
		}
		return decodeResponse(resp)
	}
}

// UpdateCurrentUser returns an endpoint that makes HTTP requests to the User
// service Update current user server.
func (c *Client) UpdateCurrentUser() goa.Endpoint {
	var (
		encodeRequest  = EncodeUpdateCurrentUserRequest(c.encoder)
		decodeResponse = DecodeUpdateCurrentUserResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildUpdateCurrentUserRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.UpdateCurrentUserDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("User", "Update current user", err)
		}
		return decodeResponse(resp)
	}
}

// DeleteCurrentUser returns an endpoint that makes HTTP requests to the User
// service Delete current user server.
func (c *Client) DeleteCurrentUser() goa.Endpoint {
	var (
		encodeRequest  = EncodeDeleteCurrentUserRequest(c.encoder)
		decodeResponse = DecodeDeleteCurrentUserResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildDeleteCurrentUserRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.DeleteCurrentUserDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("User", "Delete current user", err)
		}
		return decodeResponse(resp)
	}
}
