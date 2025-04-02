package userapi

import (
	"net/http"

	userdomain "github.com/ericprd/technical-test/internal/domain/user"
	usersvc "github.com/ericprd/technical-test/internal/service/user"
	"github.com/ericprd/technical-test/internal/utils/request"
	"github.com/ericprd/technical-test/internal/utils/response"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

func Login(
	router chi.Router,
	validate *validator.Validate,
	userSvc usersvc.User,
) {
	router.Post("/user/login", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var spec userdomain.LoginSpec

		if err := request.GetRequestBody(r, &spec); err != nil {
			response.ErrorResponse(w, err.Error(), "UNMARSHAL_BODY", http.StatusBadRequest)
			return
		}

		if err := validate.Struct(spec); err != nil {
			response.ErrorResponse(w, err.(validator.ValidationErrors).Error(), "VALIDATION_ERROR", http.StatusBadRequest)
			return
		}

		token, err := userSvc.Login(ctx, spec)
		if err != nil {
			response.ErrorResponse(w, err.Error(), "LOGIN_USER", http.StatusBadRequest)
			return
		}

		resp := map[string]string{
			"token": token,
		}

		response.SuccessResponse(w, "SUCCESS", resp, http.StatusCreated)
	})
}
