package grpc

import (
	"context"
	"github.com/garfieldlw/go-kit-demo/pages/service/demo/common"
	"github.com/garfieldlw/go-kit-demo/proto/demo"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/examples/addsvc/pkg/addendpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport/grpc"
	"os"
)

type DemoService struct {
	GetValueHandler grpc.Handler
}

func (s *DemoService) GetValue(ctx context.Context, in *demo_proto.Request) (*demo_proto.Response, error) {
	_, rsp, err := s.GetValueHandler.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}
	return rsp.(*demo_proto.Response), err
}

func GetDomeService() *DemoService {
	logger := log.NewLogfmtLogger(os.Stderr)

	d := &DemoService{
		GetValueHandler: grpc.NewServer(
			addendpoint.LoggingMiddleware(logger)(getValueEndpoint()),
			decodeRequest,
			encodeResponse,
		),
	}

	return d
}

func decodeRequest(_ context.Context, req interface{}) (interface{}, error) {
	return req, nil
}

func encodeResponse(_ context.Context, req interface{}) (interface{}, error) {
	return req, nil
}

func getValueEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*demo_proto.Request)

		resp, _ := common.GetValue(ctx, req.Rep)

		bl := new(demo_proto.Response)
		bl.Resp = resp
		return bl, nil
	}
}
