package http

import (
	"context"
	"github.com/garfieldlw/go-kit-demo/pages/service/demo/common"
	"github.com/garfieldlw/go-kit-demo/proto/demo"
	"github.com/go-kit/kit/endpoint"
)

func GetValueEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(map[string]interface{})

		resp, _ := common.GetValue(ctx, req["req"].(string))

		bl := new(demo_proto.Response)
		bl.Resp = resp
		return bl, nil
	}
}
