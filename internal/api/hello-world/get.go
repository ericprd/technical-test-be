package helloworld

import (
	"net/http"

	middlewaresvc "github.com/ericprd/technical-test/internal/service/middleware"
	"github.com/ericprd/technical-test/internal/utils/response"
	"github.com/go-chi/chi/v5"
)

func Get(
	router chi.Router,
	middleware middlewaresvc.Middleware,
) {
	router.With(middleware.Authorize()).Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		response.SuccessResponse(w, "hello world", nil, http.StatusOK)
	})
}
