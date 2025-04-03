package userapi

import (
	"net/http"

	usersvc "github.com/ericprd/technical-test/internal/service/user"
	"github.com/ericprd/technical-test/internal/utils/response"
	"github.com/go-chi/chi/v5"
)

func GetProfile(
	router chi.Router,
	userSvc usersvc.User,
) {
	router.Get("/user/{id}/profile", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id == "" {
			response.ErrorResponse(w, "user id is required", "BAD_REQUEST", http.StatusBadGateway)
			return
		}

		user, err := userSvc.GetUser(id)
		if err != nil {
			response.ErrorResponse(w, err.Error(), "BAD_GATEWAY", http.StatusBadGateway)
			return
		}

		response.SuccessResponse(w, "SUCCESS", user, http.StatusOK)
	})
}
