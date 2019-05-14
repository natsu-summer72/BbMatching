// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// User HTTP server
//
// Command:
// $ goa gen github.com/natsu-summer72/BbMatching/design

package server

import (
	"context"
	"net/http"

	user "github.com/natsu-summer72/BbMatching/gen/user"
	goa "goa.design/goa"
	goahttp "goa.design/goa/http"
)

// Server lists the User service endpoint HTTP handlers.
type Server struct {
	Mounts            []*MountPoint
	GetCurrentUser    http.Handler
	GetUser           http.Handler
	ListUser          http.Handler
	UpdateCurrentUser http.Handler
	DeleteCurrentUser http.Handler
}

// ErrorNamer is an interface implemented by generated error structs that
// exposes the name of the error as defined in the design.
type ErrorNamer interface {
	ErrorName() string
}

// MountPoint holds information about the mounted endpoints.
type MountPoint struct {
	// Method is the name of the service method served by the mounted HTTP handler.
	Method string
	// Verb is the HTTP method used to match requests to the mounted handler.
	Verb string
	// Pattern is the HTTP request path pattern used to match requests to the
	// mounted handler.
	Pattern string
}

// New instantiates HTTP handlers for all the User service endpoints.
func New(
	e *user.Endpoints,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"GetCurrentUser", "GET", "/v1/users/me"},
			{"GetUser", "GET", "/v1/users/{user_id}"},
			{"ListUser", "GET", "/v1/users"},
			{"UpdateCurrentUser", "PUT", "/v1/users"},
			{"DeleteCurrentUser", "DELETE", "/v1/users"},
		},
		GetCurrentUser:    NewGetCurrentUserHandler(e.GetCurrentUser, mux, dec, enc, eh),
		GetUser:           NewGetUserHandler(e.GetUser, mux, dec, enc, eh),
		ListUser:          NewListUserHandler(e.ListUser, mux, dec, enc, eh),
		UpdateCurrentUser: NewUpdateCurrentUserHandler(e.UpdateCurrentUser, mux, dec, enc, eh),
		DeleteCurrentUser: NewDeleteCurrentUserHandler(e.DeleteCurrentUser, mux, dec, enc, eh),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "User" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.GetCurrentUser = m(s.GetCurrentUser)
	s.GetUser = m(s.GetUser)
	s.ListUser = m(s.ListUser)
	s.UpdateCurrentUser = m(s.UpdateCurrentUser)
	s.DeleteCurrentUser = m(s.DeleteCurrentUser)
}

// Mount configures the mux to serve the User endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountGetCurrentUserHandler(mux, h.GetCurrentUser)
	MountGetUserHandler(mux, h.GetUser)
	MountListUserHandler(mux, h.ListUser)
	MountUpdateCurrentUserHandler(mux, h.UpdateCurrentUser)
	MountDeleteCurrentUserHandler(mux, h.DeleteCurrentUser)
}

// MountGetCurrentUserHandler configures the mux to serve the "User" service
// "Get current user" endpoint.
func MountGetCurrentUserHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/v1/users/me", f)
}

// NewGetCurrentUserHandler creates a HTTP handler which loads the HTTP request
// and calls the "User" service "Get current user" endpoint.
func NewGetCurrentUserHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		decodeRequest  = DecodeGetCurrentUserRequest(mux, dec)
		encodeResponse = EncodeGetCurrentUserResponse(enc)
		encodeError    = EncodeGetCurrentUserError(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "Get current user")
		ctx = context.WithValue(ctx, goa.ServiceKey, "User")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}

		res, err := endpoint(ctx, payload)

		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			eh(ctx, w, err)
		}
	})
}

// MountGetUserHandler configures the mux to serve the "User" service "Get
// User" endpoint.
func MountGetUserHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/v1/users/{user_id}", f)
}

// NewGetUserHandler creates a HTTP handler which loads the HTTP request and
// calls the "User" service "Get User" endpoint.
func NewGetUserHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		decodeRequest  = DecodeGetUserRequest(mux, dec)
		encodeResponse = EncodeGetUserResponse(enc)
		encodeError    = EncodeGetUserError(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "Get User")
		ctx = context.WithValue(ctx, goa.ServiceKey, "User")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}

		res, err := endpoint(ctx, payload)

		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			eh(ctx, w, err)
		}
	})
}

// MountListUserHandler configures the mux to serve the "User" service "List
// User" endpoint.
func MountListUserHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/v1/users", f)
}

// NewListUserHandler creates a HTTP handler which loads the HTTP request and
// calls the "User" service "List User" endpoint.
func NewListUserHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		decodeRequest  = DecodeListUserRequest(mux, dec)
		encodeResponse = EncodeListUserResponse(enc)
		encodeError    = EncodeListUserError(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "List User")
		ctx = context.WithValue(ctx, goa.ServiceKey, "User")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}

		res, err := endpoint(ctx, payload)

		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			eh(ctx, w, err)
		}
	})
}

// MountUpdateCurrentUserHandler configures the mux to serve the "User" service
// "Update current user" endpoint.
func MountUpdateCurrentUserHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("PUT", "/v1/users", f)
}

// NewUpdateCurrentUserHandler creates a HTTP handler which loads the HTTP
// request and calls the "User" service "Update current user" endpoint.
func NewUpdateCurrentUserHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		decodeRequest  = DecodeUpdateCurrentUserRequest(mux, dec)
		encodeResponse = EncodeUpdateCurrentUserResponse(enc)
		encodeError    = EncodeUpdateCurrentUserError(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "Update current user")
		ctx = context.WithValue(ctx, goa.ServiceKey, "User")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}

		res, err := endpoint(ctx, payload)

		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			eh(ctx, w, err)
		}
	})
}

// MountDeleteCurrentUserHandler configures the mux to serve the "User" service
// "Delete current user" endpoint.
func MountDeleteCurrentUserHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("DELETE", "/v1/users", f)
}

// NewDeleteCurrentUserHandler creates a HTTP handler which loads the HTTP
// request and calls the "User" service "Delete current user" endpoint.
func NewDeleteCurrentUserHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		decodeRequest  = DecodeDeleteCurrentUserRequest(mux, dec)
		encodeResponse = EncodeDeleteCurrentUserResponse(enc)
		encodeError    = EncodeDeleteCurrentUserError(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "Delete current user")
		ctx = context.WithValue(ctx, goa.ServiceKey, "User")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}

		res, err := endpoint(ctx, payload)

		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			eh(ctx, w, err)
		}
	})
}
