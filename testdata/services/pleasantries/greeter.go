package pleasantries

import (
	"github.com/pacedotdev/oto/testdata/services"
)

// GreeterService is a polite API.
// You will love it.
type GreeterService interface {
	// Greet creates a Greeting for one or more people.
	Greet(GreetRequest) GreetResponse
	// GetGreetings gets a range of saved Greetings.
	GetGreetings(GetGreetingsRequest) GetGreetingsResponse
}

// GreetRequest is the request object for GreeterService.Greet.
type GreetRequest struct {
	// Names are the names of the people to greet.
	Names []string
}

// GreetResponse is the response object for GreeterService.Greet.
type GreetResponse struct {
	// Greeting is the generated Greeting.
	Greeting Greeting
}

// GetGreetingsRequest is the request object for GreeterService.GetGreetings.
type GetGreetingsRequest struct {
	// Page describes which page of data to get.
	Page services.Page `tagtest`
}

// GetGreetingsResponse is the respponse object for GreeterService.GetGreetings.
type GetGreetingsResponse struct {
	Greetings []Greeting
}

// Greeting contains the pleasentry.
type Greeting struct {
	// Text is the message.
	Text string
}
