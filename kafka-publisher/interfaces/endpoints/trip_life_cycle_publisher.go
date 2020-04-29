package endpoints

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"Company import path goes here"
)

func MakeEventEndpoint(publisher publisher_service.Publisher) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		ev := req.(request.PayLoad)
		return publisher.Publish(ctx, ev.Value.Type, ev.Value, ev.Key)
	}
}
