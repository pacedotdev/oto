package pleasantries

import (
	"github.com/pacedotdev/oto/testdata/services"
)

type GreeterService interface {
	Greet(GreetRequest) GreetResponse
	GetGreetings(GetGreetingsRequest) GetGreetingsResponse
}

type GreetRequest struct {
	Names []string
}

type GreetResponse struct {
	Greeting Greeting
}

type GetGreetingsRequest struct {
	Page services.Page
}

type GetGreetingsResponse struct {
	Greetings []Greeting
}

type Greeting struct {
	Text string
}
