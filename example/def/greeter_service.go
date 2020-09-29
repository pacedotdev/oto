package def

// GreeterService is a polite API for greeting people.
type GreeterService interface {
	// Greet prepares a lovely greeting.
	Greet(GreetRequest) GreetResponse
}

// GreetRequest is the request object for GreeterService.Greet.
type GreetRequest struct {
	// Name is the person to greet.
	// It is required.
	Name string
}

// GreetResponse is the response object containing a
// person's greeting.
type GreetResponse struct {
	// Greeting is a nice message welcoming somebody.
	Greeting string
}
