// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// HTTP request path constructors for the User service.
//
// Command:
// $ goa gen github.com/natsu-summer72/BbMatching/design

package client

import (
	"fmt"
)

// GetCurrentUserUserPath returns the URL path to the User service Get current user HTTP endpoint.
func GetCurrentUserUserPath() string {
	return "/v1/users/me"
}

// GetUserUserPath returns the URL path to the User service Get User HTTP endpoint.
func GetUserUserPath(userID string) string {
	return fmt.Sprintf("/v1/users/%v", userID)
}

// ListUserUserPath returns the URL path to the User service List User HTTP endpoint.
func ListUserUserPath() string {
	return "/v1/users"
}

// UpdateCurrentUserUserPath returns the URL path to the User service Update current user HTTP endpoint.
func UpdateCurrentUserUserPath() string {
	return "/v1/users"
}

// DeleteCurrentUserUserPath returns the URL path to the User service Delete current user HTTP endpoint.
func DeleteCurrentUserUserPath() string {
	return "/v1/users"
}