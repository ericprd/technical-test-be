package response

import (
	"encoding/json"
	"log"
	"net/http"
)

func ErrorResponse(w http.ResponseWriter, message string, err string, statusCode int) {
	response := Response{
		Message: message,
		Error:   err,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding JSON response: %v", err)
	}
}
