package server

import (
	"fmt"
	"net/http"

	"github.com/KaffeeMaschina/http-rest-api/internal/storage/postgres"
	"github.com/go-chi/chi/v5"
)

type server struct {
	db     *postgres.DB
	router *chi.Mux
}

func NewServer(db *postgres.DB, router *chi.Mux, address string) *server {
	s := server{db: db,
		router: router}
	s.StartServer(address)
	return &s
}

func (s *server) StartServer(address string) {

	http.HandleFunc("/orders", s.showOrder)
	http.ListenAndServe(address, nil)
}

func (s *server) showOrder(w http.ResponseWriter, r *http.Request) {

	o := s.db.OrderOut()
	fmt.Fprintf(w, "%v", o)

}
