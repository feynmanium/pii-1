package services

import (
	"net/http"

	"github.com/tyndyll/pii/adapters"
)

type PiiServer struct {
	ValidateAdapter *adapters.ValidationHTTPAdapter
	Mux             *http.ServeMux
}

func (server *PiiServer) SetupRoutes() {
	if server.Mux == nil {
		server.Mux = http.NewServeMux()
	}
	server.Mux.HandleFunc(`/validate`, server.ValidateAdapter.Validate)
}

func (server *PiiServer) Start() {
	http.ListenAndServe(`:7000`, server.Mux)
}
