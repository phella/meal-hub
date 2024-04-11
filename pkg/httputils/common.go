package httputils

import (
	"encoding/json"
	"net/http"
)

// JSON responds to the provided writer with the provided body as JSON
func JSON(w http.ResponseWriter, body interface{}) error {
	raw, err := json.Marshal(body)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err = w.Write(raw); err != nil {
		return err
	}

	return nil
}
