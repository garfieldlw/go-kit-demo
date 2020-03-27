package main

import (
	demo "github.com/garfieldlw/go-kit-demo/pages/service/demo/grpc"
	"github.com/garfieldlw/go-kit-demo/proto/demo"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	ls, _ := net.Listen("tcp", ":50051")
	gs := grpc.NewServer()
	demo_proto.RegisterDemoServer(gs, demo.GetDomeService())
	log.Fatal(gs.Serve(ls))
}


