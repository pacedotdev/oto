package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/pacedotdev/oto/otohttp"
)

//go:generate ./generate.sh

// greeterService implements the generated GreeterService interface.
type greeterService struct{}

func (greeterService) Greet(ctx context.Context, r GreetRequest) (*GreetResponse, error) {
	resp := &GreetResponse{
		Greeting: fmt.Sprintf("Hello, %s.", r.Name),
	}
	return resp, nil
}

func main() {
	var greeterService greeterService
	server := otohttp.NewServer()
	RegisterGreeterService(server, greeterService)
	http.Handle("/oto/", server)
	http.Handle("/", http.FileServer(http.Dir(".")))
	fmt.Println("listening at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
