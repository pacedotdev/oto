package api

import (
	"context"
	"example/generated"
)

type GreeterService struct{}

func (GreeterService) Greet(ctx context.Context, r generated.GreetRequest) (*generated.GreetResponse, error) {
	resp := &generated.GreetResponse{
		Greeting: "Hello " + r.Name,
	}
	return resp, nil
}
