package otohttp

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/pkg/errors"
)

// Server handles oto requests.
type Server struct {
	routes   map[string]http.Handler
	NotFound http.Handler
	OnErr    func(w http.ResponseWriter, r *http.Request, err error)
}

// NewServer makes a new Server.
func NewServer() *Server {
	return &Server{
		routes: make(map[string]http.Handler),
		OnErr: func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		},
	}
}

// Register adds a handler for the specified service method.
func (s *Server) Register(service, method string, h http.Handler) {
	s.routes[fmt.Sprintf("/oto/%s.%s", service, method)] = h
}

// ServeHTTP serves the request.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		s.NotFound.ServeHTTP(w, r)
		return
	}
	h, ok := s.routes[r.URL.Path]
	if !ok {
		s.NotFound.ServeHTTP(w, r)
		return
	}
	h.ServeHTTP(w, r)
}

// Encode writes the response.
func Encode(w http.ResponseWriter, r *http.Request, status int, v interface{}) error {
	b, err := json.Marshal(v)
	if err != nil {
		return errors.Wrap(err, "encode json")
	}
	var out io.Writer = w
	if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
		w.Header().Set("Content-Encoding", "gzip")
		gzw := gzip.NewWriter(w)
		out = gzw
		defer gzw.Close()
	}
	w.Header().Set("Content-Type", "application/json; chatset=utf-8")
	w.WriteHeader(status)
	if _, err := out.Write(b); err != nil {
		return err
	}
	return nil
}

// Decode unmarshals the object in the request into v.
func Decode(r *http.Request, v interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return errors.Wrap(err, "decode json")
	}
	return nil
}
