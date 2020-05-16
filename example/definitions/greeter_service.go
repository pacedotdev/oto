package definitions

type GreeterService interface {
	Greet(GreetRequest) GreetResponse
}

type GreetRequest struct {
	Name string
}

type GreetResponse struct {
	Greeting string
}
