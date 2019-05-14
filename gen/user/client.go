// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// User client
//
// Command:
// $ goa gen github.com/natsu-summer72/BbMatching/design

package user

import (
	"context"

	goa "goa.design/goa"
)

// Client is the "User" service client.
type Client struct {
	GetCurrentUserEndpoint    goa.Endpoint
	GetUserEndpoint           goa.Endpoint
	ListUserEndpoint          goa.Endpoint
	UpdateCurrentUserEndpoint goa.Endpoint
	DeleteCurrentUserEndpoint goa.Endpoint
}

// NewClient initializes a "User" service client given the endpoints.
func NewClient(getCurrentUser, getUser, listUser, updateCurrentUser, deleteCurrentUser goa.Endpoint) *Client {
	return &Client{
		GetCurrentUserEndpoint:    getCurrentUser,
		GetUserEndpoint:           getUser,
		ListUserEndpoint:          listUser,
		UpdateCurrentUserEndpoint: updateCurrentUser,
		DeleteCurrentUserEndpoint: deleteCurrentUser,
	}
}

// GetCurrentUser calls the "Get current user" endpoint of the "User" service.
func (c *Client) GetCurrentUser(ctx context.Context, p *SessionTokenPayload) (res *BbmatchingUser, err error) {
	var ires interface{}
	ires, err = c.GetCurrentUserEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*BbmatchingUser), nil
}

// GetUser calls the "Get User" endpoint of the "User" service.
func (c *Client) GetUser(ctx context.Context, p *GetUserPayload) (res *BbmatchingUser, err error) {
	var ires interface{}
	ires, err = c.GetUserEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*BbmatchingUser), nil
}

// ListUser calls the "List User" endpoint of the "User" service.
func (c *Client) ListUser(ctx context.Context, p *SessionTokenPayload) (res BbmatchingUserCollection, err error) {
	var ires interface{}
	ires, err = c.ListUserEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(BbmatchingUserCollection), nil
}

// UpdateCurrentUser calls the "Update current user" endpoint of the "User"
// service.
func (c *Client) UpdateCurrentUser(ctx context.Context, p *UpdateUserPayload) (res *BbmatchingUser, err error) {
	var ires interface{}
	ires, err = c.UpdateCurrentUserEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*BbmatchingUser), nil
}

// DeleteCurrentUser calls the "Delete current user" endpoint of the "User"
// service.
func (c *Client) DeleteCurrentUser(ctx context.Context, p *SessionTokenPayload) (err error) {
	_, err = c.DeleteCurrentUserEndpoint(ctx, p)
	return
}
