package def

// GreeterService is a polite API for greeting people.
type GreeterService interface {
	Greet(GreetRequest) GreetResponse
}

// GreetRequest is the request object for GreeterService.Greet.
type GreetRequest struct {
	// Name is the person to greet.
	// It is required.
	Name string
}

type GreetResponse struct {
	// Greeting is a nice message welcoming somebody.
	Greeting string
}
