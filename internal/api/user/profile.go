package userapi

import (
	"net/http"

	middlewaresvc "github.com/ericprd/technical-test/internal/service/middleware"
	usersvc "github.com/ericprd/technical-test/internal/service/user"
	authutil "github.com/ericprd/technical-test/internal/utils/auth"
	jwtutil "github.com/ericprd/technical-test/internal/utils/jwt"
	"github.com/ericprd/technical-test/internal/utils/response"
	"github.com/go-chi/chi/v5"
)

func GetProfile(
	router chi.Router,
	userSvc usersvc.User,
	middleware middlewaresvc.Middleware,
) {
	router.With(middleware.Authorize()).Get("/user/profile", func(w http.ResponseWriter, r *http.Request) {
		token, err := authutil.GetTokenHeader(r)

		if err != nil {
			response.ErrorResponse(w, err.Error(), "AUTHORIZATION", http.StatusBadGateway)
			return
		}

		claims, err := jwtutil.VerifyJWT(token)
		if err != nil {
			response.ErrorResponse(w, err.Error(), "AUTHORIZATION", http.StatusBadGateway)
			return
		}

		id, ok := claims[string(jwtutil.SESSION_ID)]
		if !ok {
			response.ErrorResponse(w, "session id is not valid", "AUTHORIZATION", http.StatusBadGateway)
			return
		}

		user, err := userSvc.GetUser(id.(string))
		if err != nil {
			response.ErrorResponse(w, err.Error(), "BAD_GATEWAY", http.StatusBadGateway)
			return
		}

		response.SuccessResponse(w, "SUCCESS", user, http.StatusOK)
	})
}
