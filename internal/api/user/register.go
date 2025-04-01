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

func Register(
	router chi.Router,
	validate *validator.Validate,
	userSvc usersvc.User,
) {
	router.Post("/user/register", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var spec userdomain.Request

		if err := request.GetRequestBody(r, &spec); err != nil {
			response.ErrorResponse(w, "UNMARSHAL_BODY", err.Error(), http.StatusBadRequest)
			return
		}

		if err := validate.Struct(spec); err != nil {
			response.ErrorResponse(w, "VALIDATION_ERROR", err.(validator.ValidationErrors).Error(), http.StatusBadRequest)
			return
		}

		token, err := userSvc.Register(ctx, spec)
		if err != nil {
			response.ErrorResponse(w, "REGISTER_USER", err.Error(), http.StatusBadRequest)
			return
		}

		resp := map[string]string{
			"token": token,
		}

		response.SuccessResponse(w, "SUCCESS", resp, http.StatusCreated)
	})
}
