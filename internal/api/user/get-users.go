package userapi

import (
	"net/http"
	"time"

	"github.com/ericprd/technical-test/internal/domain/filter"
	usersvc "github.com/ericprd/technical-test/internal/service/user"
	"github.com/ericprd/technical-test/internal/utils/response"
	timeutil "github.com/ericprd/technical-test/internal/utils/time"
	"github.com/go-chi/chi/v5"
)

func GetUsers(
	router chi.Router,
	userSvc usersvc.User,
) {
	router.Get("/user/get-all", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		var (
			startDate,
			endDate time.Time

			err error
		)

		queryStartDate := q.Get(string(filter.StartDate))
		queryEndDate := q.Get(string(filter.EndDate))

		if queryStartDate != "" {
			startDate, err = timeutil.Parse(queryStartDate)
			if err != nil {
				response.ErrorResponse(w, err.Error(), "BAD_REQUEST", http.StatusBadRequest)
				return

			}
		}

		if queryEndDate != "" {
			endDate, err = timeutil.Parse(q.Get(string(filter.EndDate)))
			if err != nil {
				response.ErrorResponse(w, err.Error(), "BAD_REQUEST", http.StatusBadRequest)
				return
			}
		}

		f := filter.FilterSpec{
			Username:      q.Get(string(filter.Username)),
			AccountName:   q.Get(string(filter.AccountName)),
			AccountNumber: q.Get(string(filter.AccountNumber)),
			StartDate:     startDate,
			EndDate:       endDate,
		}

		users, err := userSvc.GetAllUser(f)
		if err != nil {
			response.ErrorResponse(w, err.Error(), "BAD_GATEWAY", http.StatusBadGateway)
			return
		}

		response.SuccessResponse(w, "SUCCESS", users, http.StatusOK)
	})
}
