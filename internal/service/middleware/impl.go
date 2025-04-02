package middlewaresvc

import (
	"net/http"

	authsvc "github.com/ericprd/technical-test/internal/service/auth"
)

type Handler func(http.Handler) http.Handler

type impl struct {
	authSvc authsvc.Auth
}

type Middleware interface {
	Authorize() Handler
}

func New(authSvc authsvc.Auth) Middleware {
	return &impl{authSvc}
}
