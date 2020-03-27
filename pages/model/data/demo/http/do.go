package http

import (
	"context"
	"github.com/garfieldlw/go-kit-demo/proto/demo"
	transport "github.com/go-kit/kit/transport/http"
	"go.elastic.co/apm"
	"go.elastic.co/apm/module/apmhttp"
	"net/http"
	"net/url"
)

func GetValue(ctx context.Context, req *demo_proto.Request) (*demo_proto.Response, error) {
	client := transport.NewClient(
		"GET", &url.URL{ /*...*/ },
		transport.EncodeJSONRequest,
		func(_ context.Context, r *http.Response) (interface{}, error) { return nil, nil },
		transport.SetClient(apmhttp.WrapClient(http.DefaultClient)),
	)

	tx := apm.DefaultTracer.StartTransaction("name", "type")
	ctx = apm.ContextWithTransaction(ctx, tx)
	defer tx.End()

	_, err := client.Endpoint()(ctx, struct{}{})
	if err != nil {
		return nil, err
	}

	return nil, nil
}
