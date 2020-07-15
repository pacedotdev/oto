package pleasantries

import (
	"github.com/pacedotdev/oto/testdata/services"
)

// GreeterService is a polite API.
type GreeterService interface {
	// Greet creates a Greeting for one or more people.
	Greet(GreetRequest) GreetResponse
	// GetGreetings gets a range of saved Greetings.
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

// Greeting contains the pleasentry.
type Greeting struct {
	// Text is the message.
	Text string
}
