package vault

import (
	httptransport "github.com/go-kit/kit/transport/http"
	"golang.org/x/net/context"
	"net/http"
)

// NewHTTPServer makes a new Vault HTTP service.
func NewHTTPServer(ctx context.Context, endpoints Endpoints) http.Handler {
	m := http.NewServeMux()
	m.Handle("/hash", httptransport.NewServer(

		endpoints.HashEndpoint,
		decodeHashRequest,
		encodeResponse,
	))
	m.Handle("/validate", httptransport.NewServer(

		endpoints.ValidateEndpoint,
		decodeValidateRequest,
		encodeResponse,
	))
	return m
}
