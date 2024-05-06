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

func NewServer(db *postgres.DB, router *chi.Mux) *server {
	s := server{db: db,
		router: router}
	s.StartServer()
	return &s
}

func (s *server) StartServer() {

	http.HandleFunc("/orders", s.showOrder)
	http.ListenAndServe(":8082", nil)
}

func (s *server) showOrder(w http.ResponseWriter, r *http.Request) {
	fmt.Println("f[f[f]]")
	o := s.db.OrderOut()
	fmt.Fprintf(w, "%v", o)

}
