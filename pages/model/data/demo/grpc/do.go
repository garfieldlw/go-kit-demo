package grpc

import (
	"context"
	"flag"
	"github.com/garfieldlw/go-kit-demo/proto/demo"
	transport "github.com/go-kit/kit/transport/grpc"
	"go.elastic.co/apm"
	"go.elastic.co/apm/module/apmgrpc"
	"google.golang.org/grpc"
	"log"
	"time"
)

func GetValue(ctx context.Context, req *demo_proto.Request) (*demo_proto.Response, error) {
	addr := flag.String("addr", ":8081", "gRPC address")

	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithUnaryInterceptor(apmgrpc.NewUnaryClientInterceptor()), grpc.WithTimeout(1*time.Second))

	if err != nil {
		log.Fatalln("gRPC dial:", err)
	}
	defer conn.Close()

	tx := apm.DefaultTracer.StartTransaction("name", "type")
	ctx = apm.ContextWithTransaction(ctx, tx)
	defer tx.End()

	resp, err := transport.NewClient(
		conn, "demo", "demo",
		decodeRequest,
		encodeResponse,
		demo_proto.Response{},
	).Endpoint()(ctx, req)

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
