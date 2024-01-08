// Code generated by goa v3.14.2, DO NOT EDIT.
//
// front client HTTP transport
//
// Command:
// $ goa gen goa.design/clue/example/weather/services/front/design -o
// services/front

package client

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the front service endpoint HTTP clients.
type Client struct {
	// Forecast Doer is the HTTP client used to make requests to the forecast
	// endpoint.
	ForecastDoer goahttp.Doer

	// TestAll Doer is the HTTP client used to make requests to the test_all
	// endpoint.
	TestAllDoer goahttp.Doer

	// TestSmoke Doer is the HTTP client used to make requests to the test_smoke
	// endpoint.
	TestSmokeDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the front service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		ForecastDoer:        doer,
		TestAllDoer:         doer,
		TestSmokeDoer:       doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// Forecast returns an endpoint that makes HTTP requests to the front service
// forecast server.
func (c *Client) Forecast() goa.Endpoint {
	var (
		decodeResponse = DecodeForecastResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildForecastRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ForecastDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("front", "forecast", err)
		}
		return decodeResponse(resp)
	}
}

// TestAll returns an endpoint that makes HTTP requests to the front service
// test_all server.
func (c *Client) TestAll() goa.Endpoint {
	var (
		encodeRequest  = EncodeTestAllRequest(c.encoder)
		decodeResponse = DecodeTestAllResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildTestAllRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.TestAllDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("front", "test_all", err)
		}
		return decodeResponse(resp)
	}
}

// TestSmoke returns an endpoint that makes HTTP requests to the front service
// test_smoke server.
func (c *Client) TestSmoke() goa.Endpoint {
	var (
		decodeResponse = DecodeTestSmokeResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildTestSmokeRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.TestSmokeDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("front", "test_smoke", err)
		}
		return decodeResponse(resp)
	}
}
