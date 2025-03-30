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

	// Set response headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	// Encode and send the response as JSON
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding JSON response: %v", err)
	}
}
