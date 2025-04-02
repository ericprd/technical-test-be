package userapi

import (
	"fmt"
	"net/http"

	bankdomain "github.com/ericprd/technical-test/internal/domain/bank"
	banksvc "github.com/ericprd/technical-test/internal/service/bank"
	middlewaresvc "github.com/ericprd/technical-test/internal/service/middleware"
	authutil "github.com/ericprd/technical-test/internal/utils/auth"
	jwtutil "github.com/ericprd/technical-test/internal/utils/jwt"
	"github.com/ericprd/technical-test/internal/utils/request"
	"github.com/ericprd/technical-test/internal/utils/response"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

func CreateBank(
	router chi.Router,
	validate *validator.Validate,
	bankSvc banksvc.Bank,
	middleware middlewaresvc.Middleware,
) {
	router.With(middleware.Authorize()).Post("/user/create-bank", func(w http.ResponseWriter, r *http.Request) {
		token, err := authutil.GetTokenHeader(r)
		if err != nil {
			response.ErrorResponse(w, err.Error(), "UNAUTHORIZED", http.StatusUnauthorized)
			return
		}

		claims, err := jwtutil.VerifyJWT(token)
		if err != nil {
			response.ErrorResponse(w, err.Error(), "UNAUTHORIZED", http.StatusUnauthorized)
			return
		}

		sessionID, ok := claims[string(jwtutil.SESSION_ID)]
		if !ok {
			response.ErrorResponse(w, "invalid session id", "UNAUTHORIZED", http.StatusUnauthorized)
			return
		}

		var spec bankdomain.RegisterSpec

		if err := request.GetRequestBody(r, &spec); err != nil {
			response.ErrorResponse(w, err.Error(), "UNMARSHAL_BODY", http.StatusBadRequest)
			return
		}

		if err := validate.Struct(spec); err != nil {
			response.ErrorResponse(w, err.(validator.ValidationErrors).Error(), "VALIDATION_ERROR", http.StatusBadRequest)
			return
		}

		spec.UserID = fmt.Sprintf("%v", sessionID)

		if err := bankSvc.Register(spec); err != nil {
			response.ErrorResponse(w, err.Error(), "BAD_REQUEST", http.StatusBadRequest)
			return
		}

		response.SuccessResponse(w, "SUCCESS", map[string]string{
			"message": "user's bank created succesfully",
		}, http.StatusOK)
	})
}
