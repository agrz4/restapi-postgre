package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/agrz4/rest-api-postgre/services/users"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()

	// register services
	userStore := users.NewUserStore(s.db)
	userHandler := users.NewHandler(userStore)
	userHandler.RegisterRoutes(router)

	log.Println("Listening on port", s.addr)

	return http.ListenAndServe(s.addr, router)
}
