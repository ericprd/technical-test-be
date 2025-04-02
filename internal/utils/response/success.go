package response

import (
	"encoding/json"
	"log"
	"net/http"
)

func SuccessResponse(w http.ResponseWriter, message string, data any, statusCode int) {
	response := Response{
		Message: message,
		Data:    data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding JSON response: %v", err)
	}
}
