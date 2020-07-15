// Code generated by oto; DO NOT EDIT.

package main

import (
	"context"
	"log"
	"net/http"

	"github.com/pacedotdev/oto/otohttp"
)

// GreeterService is a polite API for greeting people.
type GreeterService interface {
	Greet(context.Context, GreetRequest) (*GreetResponse, error)
}

type greeterServiceServer struct {
	server         *otohttp.Server
	greeterService GreeterService
}

// Register adds the GreeterService to the otohttp.Server.
func RegisterGreeterService(server *otohttp.Server, greeterService GreeterService) {
	handler := &greeterServiceServer{
		server:         server,
		greeterService: greeterService,
	}
	server.Register("GreeterService", "Greet", handler.handleGreet)
}

func (s *greeterServiceServer) handleGreet(w http.ResponseWriter, r *http.Request) {
	var request GreetRequest
	if err := otohttp.Decode(r, &request); err != nil {
		s.server.OnErr(w, r, err)
		return
	}
	response, err := s.greeterService.Greet(r.Context(), request)
	if err != nil {
		log.Println("TODO: oto service error:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := otohttp.Encode(w, r, http.StatusOK, response); err != nil {
		s.server.OnErr(w, r, err)
		return
	}
}

// GreetRequest is the request object for GreeterService.Greet.
type GreetRequest struct {
	// Name is the person to greet. It is required.
	Name string `json:"name"`
}

// GreetResponse is the response object for GreeterService.Greet.
type GreetResponse struct {
	// Greeting is a nice message welcoming somebody.
	Greeting string `json:"greeting"`
	// Error is string explaining what went wrong. Empty if everything was fine.
	Error string `json:"error,omitempty"`
}
