// Code generated by goa v3.5.4, DO NOT EDIT.
//
// Forecaster gRPC client
//
// Command:
// $ goa gen
// github.com/goadesign/clue/example/weather/services/forecaster/design -o
// services/forecaster

package client

import (
	"context"

	forecasterpb "github.com/goadesign/clue/example/weather/services/forecaster/gen/grpc/forecaster/pb"
	goagrpc "goa.design/goa/v3/grpc"
	goa "goa.design/goa/v3/pkg"
	"google.golang.org/grpc"
)

// Client lists the service endpoint gRPC clients.
type Client struct {
	grpccli forecasterpb.ForecasterClient
	opts    []grpc.CallOption
}

// NewClient instantiates gRPC client for all the Forecaster service servers.
func NewClient(cc *grpc.ClientConn, opts ...grpc.CallOption) *Client {
	return &Client{
		grpccli: forecasterpb.NewForecasterClient(cc),
		opts:    opts,
	}
}

// Forecast calls the "Forecast" function in forecasterpb.ForecasterClient
// interface.
func (c *Client) Forecast() goa.Endpoint {
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		inv := goagrpc.NewInvoker(
			BuildForecastFunc(c.grpccli, c.opts...),
			EncodeForecastRequest,
			DecodeForecastResponse)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			return nil, goa.Fault(err.Error())
		}
		return res, nil
	}
}
