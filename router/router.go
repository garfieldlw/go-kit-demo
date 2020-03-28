package router

import (
	"context"
	"encoding/json"
	http2 "github.com/garfieldlw/go-kit-demo/pages/service/demo/http"
	"github.com/go-kit/kit/examples/addsvc/pkg/addendpoint"
	"github.com/go-kit/kit/log"
	transport "github.com/go-kit/kit/transport/http"
	"go.elastic.co/apm/module/apmhttp"
	"net/http"
	"os"
)

func LoadRouter() {
	logger := log.NewLogfmtLogger(os.Stderr)

	http.Handle("/demo", apmhttp.Wrap(transport.NewServer(
		addendpoint.LoggingMiddleware(logger)( http2.GetValueEndpoint()),
		decodeRequest,
		encodeResponse,
	)))
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
