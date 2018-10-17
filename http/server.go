package http

import (
	"github.com/codegangsta/negroni"
)

// Server represents an HTTP server.
type Server struct {
	*negroni.Negroni
}

// NewServer returns a new instance of Server.
func NewServer() *Server {
	return &Server{negroni.Classic()}
}
