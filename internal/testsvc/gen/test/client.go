// Code generated by goa v3.5.4, DO NOT EDIT.
//
// test client
//
// Command:
// $ goa gen github.com/goadesign/clue/internal/testsvc/design

package test

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "test" service client.
type Client struct {
	HTTPMethodEndpoint goa.Endpoint
	GrpcMethodEndpoint goa.Endpoint
	GrpcStreamEndpoint goa.Endpoint
}

// NewClient initializes a "test" service client given the endpoints.
func NewClient(hTTPMethod, grpcMethod, grpcStream goa.Endpoint) *Client {
	return &Client{
		HTTPMethodEndpoint: hTTPMethod,
		GrpcMethodEndpoint: grpcMethod,
		GrpcStreamEndpoint: grpcStream,
	}
}

// HTTPMethod calls the "http_method" endpoint of the "test" service.
func (c *Client) HTTPMethod(ctx context.Context, p *Fields) (res *Fields, err error) {
	var ires interface{}
	ires, err = c.HTTPMethodEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*Fields), nil
}

// GrpcMethod calls the "grpc_method" endpoint of the "test" service.
func (c *Client) GrpcMethod(ctx context.Context, p *Fields) (res *Fields, err error) {
	var ires interface{}
	ires, err = c.GrpcMethodEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*Fields), nil
}

// GrpcStream calls the "grpc_stream" endpoint of the "test" service.
func (c *Client) GrpcStream(ctx context.Context) (res GrpcStreamClientStream, err error) {
	var ires interface{}
	ires, err = c.GrpcStreamEndpoint(ctx, nil)
	if err != nil {
		return
	}
	return ires.(GrpcStreamClientStream), nil
}
