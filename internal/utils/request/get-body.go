package request

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetRequestBody(r *http.Request, target any) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return fmt.Errorf("error reading request body: %v", err)
	}

	defer r.Body.Close()

	if err := json.Unmarshal(body, target); err != nil {
		return fmt.Errorf("error unmarshalling request body: %v", err)
	}

	return nil
}
