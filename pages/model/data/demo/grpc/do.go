package grpc

import (
	"context"
	"flag"
	"github.com/garfieldlw/go-kit-demo/proto/demo"
	"github.com/go-kit/kit/examples/addsvc/pkg/addendpoint"
	"github.com/go-kit/kit/log"
	transport "github.com/go-kit/kit/transport/grpc"
	"go.elastic.co/apm"
	"go.elastic.co/apm/module/apmgrpc"
	"google.golang.org/grpc"
	"os"
	"time"
)

func GetValue(ctx context.Context, req *demo_proto.Request) (*demo_proto.Response, error) {
	addr := flag.String("addr", ":8081", "gRPC address")

	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithUnaryInterceptor(apmgrpc.NewUnaryClientInterceptor()), grpc.WithTimeout(1*time.Second))

	if err != nil {
		return nil, err
	}
	defer conn.Close()

	tx := apm.DefaultTracer.StartTransaction("name", "type")
	ctx = apm.ContextWithTransaction(ctx, tx)
	defer tx.End()

	ep := transport.NewClient(
		conn, "demo", "demo",
		decodeRequest,
		encodeResponse,
		demo_proto.Response{},
	).Endpoint()

	logger := log.NewLogfmtLogger(os.Stderr)
	ep = addendpoint.LoggingMiddleware(logger)(ep)

	resp, err := ep(ctx, req)

	if err != nil {
		return nil, err
	}

	return resp.(*demo_proto.Response), nil
}

func decodeRequest(_ context.Context, req interface{}) (interface{}, error) {
	return req, nil
}

func encodeResponse(_ context.Context, req interface{}) (interface{}, error) {
	return req, nil
}
