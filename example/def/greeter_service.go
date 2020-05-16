package def

// GreeterService is a polite API for greeting people.
type GreeterService interface {
	Greet(GreetRequest) GreetResponse
}

type GreetRequest struct {
	// Name is the person to greet.
	Name string
}

type GreetResponse struct {
	// Greeting is a nice message welcoming somebody.
	Greeting string
}
