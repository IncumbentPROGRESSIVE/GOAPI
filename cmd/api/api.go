package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)
type APIserver struct {
	addr string 
	db *sql.DB
}

func NewAPIsever(addr string, db *sql.DB) *APIserver{
	return &APIserver{
		addr: addr,
		db: db,
	}
}

func (s *APIserver) start() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()


	log.Println("listening on", s.addr)
	return http.ListenAndServe(s.addr, router)
}