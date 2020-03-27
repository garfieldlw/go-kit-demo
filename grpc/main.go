package main

import (
	"github.com/garfieldlw/go-kit-demo/pages/service"
	"github.com/garfieldlw/go-kit-demo/proto/demo"
	"google.golang.org/grpc"
	"net"
)

func main() {
	ls, _ := net.Listen("tcp", ":50051")
	gs := grpc.NewServer()
	demo_proto.RegisterDemoServer(gs, service.GetDomeService())
	gs.Serve(ls)
}


