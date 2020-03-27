package http

import (
	"context"
	"flag"
	"github.com/garfieldlw/go-kit-demo/proto/demo"
	transport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
	"log"
	"time"
)

func GetValue(ctx context.Context, req *demo_proto.Request) (resp *demo_proto.Response, err error) {
	var (
		grpcAddr = flag.String("addr", ":8081",
			"gRPC address")
	)
	flag.Parse()
	conn, err := grpc.Dial(*grpcAddr, grpc.WithInsecure(),
		grpc.WithTimeout(1*time.Second))

	if err != nil {
		log.Fatalln("gRPC dial:", err)
	}
	defer conn.Close()

	var loremEndpoint = transport.NewClient(
		conn, "Lorem", "Lorem",
		lorem_grpc.EncodeGRPCLoremRequest,
		lorem_grpc.DecodeGRPCLoremResponse,
		pb.LoremResponse{},
	).Endpoint()

	return lorem_grpc.Endpoints{
		LoremEndpoint: loremEndpoint,
	}
}
