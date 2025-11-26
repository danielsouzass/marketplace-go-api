package jsonutils

import (
	"encoding/json"
	"fmt"
	"marketplace/internal/types"
	"net/http"
)

func SendJSON(w http.ResponseWriter, resp types.Response) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.Code)
	if resp.Body != nil {
		if err := json.NewEncoder(w).Encode(resp.Body); err != nil {
			return fmt.Errorf("failed to encode json %w", err)
		}
	}
	return nil
}

func ReadJSON[T any](r *http.Request) (T, error) {
	var data T
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		return data, fmt.Errorf("decode json failed: %w", err)
	}
	return data, nil
}
