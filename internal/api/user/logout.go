package userapi

import (
	"net/http"

	"github.com/ericprd/technical-test/internal/service/middlewaresvc"
	usersvc "github.com/ericprd/technical-test/internal/service/user"
	authutil "github.com/ericprd/technical-test/internal/utils/auth"
	"github.com/ericprd/technical-test/internal/utils/response"
	"github.com/go-chi/chi/v5"
)

func Logout(
	middleware middlewaresvc.Middleware,
	router chi.Router,
	userSvc usersvc.User,
) {
	router.With(middleware.Authorize()).Post("/user/logout", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		token, err := authutil.GetTokenHeader(r)
		if err != nil {
			response.ErrorResponse(w, err.Error(), "UNAUTHORIZED", http.StatusUnauthorized)
			return
		}

		if err := userSvc.Logout(ctx, token); err != nil {
			response.ErrorResponse(w, err.Error(), "LOGOUT", http.StatusBadRequest)
			return
		}

		response.SuccessResponse(w, "SUCCESS", map[string]string{
			"message": "logout success",
		}, http.StatusOK)
	})
}
