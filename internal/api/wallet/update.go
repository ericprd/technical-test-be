package walletapi

import (
	"net/http"

	walletdomain "github.com/ericprd/technical-test/internal/domain/wallet"
	middlewaresvc "github.com/ericprd/technical-test/internal/service/middleware"
	walletsvc "github.com/ericprd/technical-test/internal/service/wallet"
	authutil "github.com/ericprd/technical-test/internal/utils/auth"
	jwtutil "github.com/ericprd/technical-test/internal/utils/jwt"
	"github.com/ericprd/technical-test/internal/utils/request"
	"github.com/ericprd/technical-test/internal/utils/response"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

func Update(
	router chi.Router,
	validate *validator.Validate,
	walletSvc walletsvc.Wallet,
	middlewareSvc middlewaresvc.Middleware,
) {
	router.With(middlewareSvc.Authorize()).Put("/wallet/update-balance", func(w http.ResponseWriter, r *http.Request) {
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

		var spec walletdomain.UpdateSpec

		if err := request.GetRequestBody(r, &spec); err != nil {
			response.ErrorResponse(w, err.Error(), "UNMARSHAL_BODY", http.StatusBadRequest)
			return
		}

		if err := validate.Struct(spec); err != nil {
			response.ErrorResponse(w, err.(validator.ValidationErrors).Error(), "VALIDATION_ERROR", http.StatusBadRequest)
			return
		}

		spec.UserID = sessionID.(string)

		if err := walletSvc.Update(spec); err != nil {
			response.ErrorResponse(w, err.Error(), "UPDATE_BALANCE", http.StatusBadRequest)
			return
		}

		response.SuccessResponse(w, "SUCCESS", map[string]string{
			"message": "wallet balance updated succesfully",
		}, http.StatusOK)
	})
}
