package middlewaresvc

import (
	"log"
	"net/http"

	"github.com/ericprd/technical-test/internal/utils/response"
)

func (i *impl) Authorize() Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if err := i.authSvc.Verify(r); err != nil {
				response.ErrorResponse(w, err.Error(), "UNAUTHORIZE", http.StatusUnauthorized)
				return
			}
			log.Print("lewat sini")
			h.ServeHTTP(w, r.WithContext(r.Context()))
		})
	}
}
