package helloworld

import (
	"net/http"

	"github.com/ericprd/technical-test/internal/utils/response"
	"github.com/go-chi/chi/v5"
)

func Get(router chi.Router) {
	router.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		response.SuccessResponse(w, "hello world", nil, http.StatusOK)
	})
}
