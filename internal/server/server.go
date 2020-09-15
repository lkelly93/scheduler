package server

import (
	"context"
	"fmt"
	"net/http"
)

type HTTPServer interface {
	Start(port int) error

	Stop() error
}

func NewHTTPServer() HTTPServer {
	return &httpServer{}
}

type httpServer struct {
	server *http.Server
}

func (server *httpServer) Start(port int) error {
	serveMux := newServeMux()
	server.server = &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: serveMux,
	}
	return server.server.ListenAndServe()
}

func (server *httpServer) Stop() error {
	return server.server.Shutdown(context.Background())
}
