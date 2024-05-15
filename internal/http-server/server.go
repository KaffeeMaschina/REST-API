package server

import (
	"fmt"
	"net/http"

	"github.com/KaffeeMaschina/http-rest-api/internal/storage"
	"github.com/KaffeeMaschina/http-rest-api/internal/storage/postgres"
	"github.com/go-chi/chi/v5"
)

type OrderGetter interface {
	OrderOut(oid string) (o storage.Orders)
}

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
	//s.router.Get("/", s.showAllOrder)
	s.router.Get("/{oid}", s.showAllOrderByID)
	http.ListenAndServe(address, s.router)
}

func (s *server) showAllOrderByID(w http.ResponseWriter, r *http.Request) {
	oid := chi.URLParam(r, "oid")

	OrderOut := s.db.OrderOut(oid)
	fmt.Fprintf(w, "%v", OrderOut)

}

/*func (s *server) showAllOrder(w http.ResponseWriter, r *http.Request) {
	OrderOut := s.db.OrderOut(uid)
	t, err := template.ParseFiles("internal/http-server/html/order.html")
	if err != nil {
		log.Printf("getOrder(): ошибка парсинга шаблона html: %s\n", err)
		http.Error(w, "Internal Server Error", 500)
		return
	}

	t.Execute(w, OrderOut)
}*/
