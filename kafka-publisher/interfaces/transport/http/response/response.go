package response

import (
	"context"
	"encoding/json"
	"net/http"
)

type EventResponse struct {
	V interface{} `json:"data,omitempty"`
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	res := EventResponse{
		V: response,
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(res)
}
