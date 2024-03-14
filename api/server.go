package api

import (
	"net/http"
	"pocket-link/config"
	db "pocket-link/db/sqlc"
)

type Server struct {
	config config.Config
	store  db.Store
	mux    *http.ServeMux
}

func NewServer(config config.Config, store db.Store) (*Server, error) {
	server := &Server{
		config: config,
		store:  store,
	}

	if err := server.setupServeMux(); err != nil {
		return nil, err
	}
	return server, nil
}

func (server *Server) setupServeMux() error {
	mux := http.NewServeMux()
	mux.Handle("POST /shortener", http.HandlerFunc(server.CreateUrl))

	server.mux = mux
	return nil
}

func (server *Server) Start() error {
	return http.ListenAndServe(":8080", server.mux)
}
