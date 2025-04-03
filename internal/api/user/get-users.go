package userapi

import (
	"net/http"

	usersvc "github.com/ericprd/technical-test/internal/service/user"
	"github.com/ericprd/technical-test/internal/utils/response"
	"github.com/go-chi/chi/v5"
)

func GetUsers(
	router chi.Router,
	userSvc usersvc.User,
) {
	router.Get("/user/get-all", func(w http.ResponseWriter, r *http.Request) {
		users, err := userSvc.GetAllUser()
		if err != nil {
			response.ErrorResponse(w, err.Error(), "BAD_GATEWAY", http.StatusBadGateway)
			return
		}

		response.SuccessResponse(w, "SUCCESS", users, http.StatusOK)
	})
}
