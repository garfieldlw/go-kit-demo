package router

import (
	"context"
	"encoding/json"
	http2 "github.com/garfieldlw/go-kit-demo/pages/service/demo/http"
	transport "github.com/go-kit/kit/transport/http"
	"net/http"
)

func LoadRouter() {
	http.Handle("/demo", transport.NewServer(
		http2.GetValueEndpoint(),
		decodeRequest,
		encodeResponse,
	))
}

func decodeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}