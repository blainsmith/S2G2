// Code generated by oto; DO NOT EDIT.

package rpcserver

import (
	"context"
	"net/http"

	"github.com/pacedotdev/oto/otohttp"
)

// GreeterService makes nice greetings.
type GreeterService interface {

	// Greet makes a greeting.
	Greet(context.Context, GreetRequest) (*GreetResponse, error)
}

type MathService interface {
	Add(context.Context, AddRequest) (*AddResponse, error)
	Subtract(context.Context, SubtractRequest) (*SubtractResponse, error)
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
		s.server.OnErr(w, r, err)
		return
	}
	if err := otohttp.Encode(w, r, http.StatusOK, response); err != nil {
		s.server.OnErr(w, r, err)
		return
	}
}

type mathServiceServer struct {
	server      *otohttp.Server
	mathService MathService
}

// Register adds the MathService to the otohttp.Server.
func RegisterMathService(server *otohttp.Server, mathService MathService) {
	handler := &mathServiceServer{
		server:      server,
		mathService: mathService,
	}
	server.Register("MathService", "Add", handler.handleAdd)
	server.Register("MathService", "Subtract", handler.handleSubtract)
}

func (s *mathServiceServer) handleAdd(w http.ResponseWriter, r *http.Request) {
	var request AddRequest
	if err := otohttp.Decode(r, &request); err != nil {
		s.server.OnErr(w, r, err)
		return
	}
	response, err := s.mathService.Add(r.Context(), request)
	if err != nil {
		s.server.OnErr(w, r, err)
		return
	}
	if err := otohttp.Encode(w, r, http.StatusOK, response); err != nil {
		s.server.OnErr(w, r, err)
		return
	}
}

func (s *mathServiceServer) handleSubtract(w http.ResponseWriter, r *http.Request) {
	var request SubtractRequest
	if err := otohttp.Decode(r, &request); err != nil {
		s.server.OnErr(w, r, err)
		return
	}
	response, err := s.mathService.Subtract(r.Context(), request)
	if err != nil {
		s.server.OnErr(w, r, err)
		return
	}
	if err := otohttp.Encode(w, r, http.StatusOK, response); err != nil {
		s.server.OnErr(w, r, err)
		return
	}
}

type AddRequest struct {
	A float64 `json:"a"`
	B float64 `json:"b"`
}

type AddResponse struct {
	Result float64 `json:"result"`
	// Error is string explaining what went wrong. Empty if everything was fine.
	Error string `json:"error,omitempty"`
}

// GreetRequest is the request object for GreeterService.Greet.
type GreetRequest struct {
	// Name is the person to greet.
	Name string `json:"name"`
}

// GreetResponse is the response object containing a person's greeting.
type GreetResponse struct {
	// Greeting is the greeting that was generated.
	Greeting string `json:"greeting"`
	// Error is string explaining what went wrong. Empty if everything was fine.
	Error string `json:"error,omitempty"`
}

type SubtractRequest struct {
	A float64 `json:"a"`
	B float64 `json:"b"`
}

type SubtractResponse struct {
	Result float64 `json:"result"`
	// Error is string explaining what went wrong. Empty if everything was fine.
	Error string `json:"error,omitempty"`
}