// Code generated by goa v3.5.4, DO NOT EDIT.
//
// Forecaster gRPC server types
//
// Command:
// $ goa gen
// github.com/goadesign/clue/example/weather/services/forecaster/design -o
// services/forecaster

package server

import (
	forecaster "github.com/goadesign/clue/example/weather/services/forecaster/gen/forecaster"
	forecasterpb "github.com/goadesign/clue/example/weather/services/forecaster/gen/grpc/forecaster/pb"
)

// NewForecastPayload builds the payload of the "forecast" endpoint of the
// "Forecaster" service from the gRPC request type.
func NewForecastPayload(message *forecasterpb.ForecastRequest) *forecaster.ForecastPayload {
	v := &forecaster.ForecastPayload{
		Lat:  message.Lat,
		Long: message.Long,
	}
	return v
}

// NewForecastResponse builds the gRPC response type from the result of the
// "forecast" endpoint of the "Forecaster" service.
func NewForecastResponse(result *forecaster.Forecast2) *forecasterpb.ForecastResponse {
	message := &forecasterpb.ForecastResponse{}
	if result.Location != nil {
		message.Location = svcForecasterLocationToForecasterpbLocation(result.Location)
	}
	if result.Periods != nil {
		message.Periods = make([]*forecasterpb.Period, len(result.Periods))
		for i, val := range result.Periods {
			message.Periods[i] = &forecasterpb.Period{
				Name:            val.Name,
				StartTime:       val.StartTime,
				EndTime:         val.EndTime,
				Temperature:     int32(val.Temperature),
				TemperatureUnit: val.TemperatureUnit,
				Summary:         val.Summary,
			}
		}
	}
	return message
}

// svcForecasterLocationToForecasterpbLocation builds a value of type
// *forecasterpb.Location from a value of type *forecaster.Location.
func svcForecasterLocationToForecasterpbLocation(v *forecaster.Location) *forecasterpb.Location {
	res := &forecasterpb.Location{
		Lat:   v.Lat,
		Long:  v.Long,
		City:  v.City,
		State: v.State,
	}

	return res
}

// protobufForecasterpbLocationToForecasterLocation builds a value of type
// *forecaster.Location from a value of type *forecasterpb.Location.
func protobufForecasterpbLocationToForecasterLocation(v *forecasterpb.Location) *forecaster.Location {
	res := &forecaster.Location{
		Lat:   v.Lat,
		Long:  v.Long,
		City:  v.City,
		State: v.State,
	}

	return res
}
