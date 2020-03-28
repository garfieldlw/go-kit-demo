package http

import (
	"context"
	"github.com/garfieldlw/go-kit-demo/proto/demo"
	"github.com/go-kit/kit/examples/addsvc/pkg/addendpoint"
	"github.com/go-kit/kit/log"
	transport "github.com/go-kit/kit/transport/http"
	"go.elastic.co/apm"
	"go.elastic.co/apm/module/apmhttp"
	"net/http"
	"net/url"
	"os"
)

func GetValue(ctx context.Context, req *demo_proto.Request) (*demo_proto.Response, error) {
	client := transport.NewClient(
		"POST", &url.URL{ /*...*/ },
		transport.EncodeJSONRequest,
		func(_ context.Context, r *http.Response) (interface{}, error) { return nil, nil },
		transport.SetClient(apmhttp.WrapClient(http.DefaultClient)),
	).Endpoint()

	logger := log.NewLogfmtLogger(os.Stderr)
	client = addendpoint.LoggingMiddleware(logger)(client)

	tx := apm.DefaultTracer.StartTransaction("name", "type")
	ctx = apm.ContextWithTransaction(ctx, tx)
	defer tx.End()

	resp, err := client(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.(*demo_proto.Response), nil
}
