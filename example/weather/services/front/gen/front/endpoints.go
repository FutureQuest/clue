// Code generated by goa v3.5.4, DO NOT EDIT.
//
// front endpoints
//
// Command:
// $ goa gen github.com/goadesign/clue/example/weather/services/front/design
// -o services/front

package front

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "front" service endpoints.
type Endpoints struct {
	Forecast goa.Endpoint
}

// NewEndpoints wraps the methods of the "front" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		Forecast: NewForecastEndpoint(s),
	}
}

// Use applies the given middleware to all the "front" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Forecast = m(e.Forecast)
}

// NewForecastEndpoint returns an endpoint function that calls the method
// "forecast" of service "front".
func NewForecastEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(string)
		return s.Forecast(ctx, p)
	}
}
