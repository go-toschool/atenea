package session

import (
	"net/http"

	"github.com/go-toschool/palermo/auth"
	"github.com/go-toschool/syracuse/citizens"
)

// ContentTypeGraphQL graphql content type.
const (
	ContentTypeGraphQL   = "application/json"
	authUserIDContextKey = "fromTokenSessionUserId"
)

// GraphRequest struct to unmarshal query.
type GraphRequest struct {
	Query string `json:"query"`
}

// Context ...
type Context struct {
	User    citizens.CitizenshipClient
	Session auth.AuthServiceClient
}

// Handle creates a new bounded Handler with context.
func (c *Context) Handle(h HandlerFunc) *Handler {
	return &Handler{c, h}
}

// HandlerFunc function handler signature used by sigiriya application.
type HandlerFunc func(*Context, http.ResponseWriter, *http.Request)

// Handler ...
type Handler struct {
	ctx    *Context
	handle HandlerFunc
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.handle(h.ctx, w, r)
}
