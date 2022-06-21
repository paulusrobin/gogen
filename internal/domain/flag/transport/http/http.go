package http

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	kit "github.com/go-kit/kit/transport/http"
	"github.com/paulusrobin/gogen/internal/pkg/json"
	"github.com/paulusrobin/gogen/internal/pkg/validator"
	"net/http"
)

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func errorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if vErr, ok := validator.IsValidationError(err); ok {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"message": vErr.Error(),
			"errors":  vErr.Detail(),
		})
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"message": err.Error(),
			"errors":  []string{err.Error()},
		})
	}
}

// MakeHandler creates a single handler for all incoming http requests.
// Each http request participant need to pass own endpoint.Endpoint and
// kit.DecodeRequestFunc which is responsible for converting payload.
func MakeHandler(endpoint endpoint.Endpoint, d kit.DecodeRequestFunc) http.Handler {
	return kit.NewServer(
		endpoint,
		d,
		encodeResponse,
		kit.ServerErrorEncoder(errorEncoder),
		kit.ServerBefore(),
	)
}
