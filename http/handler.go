package http

import (
	"fmt"
	"net/http"

	"github.com/HugoDrl/ObservableBroker.git/storage"
)

type Server struct {
	db *storage.Data
}

func (s *Server) handleHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}

func (s *Server) handleMessages(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Request received by %v\n", r.Host)
	messages := fmt.Sprintf("%v", s.db.Messages)
	w.Write([]byte(messages))
}

func (s *Server) handleClients(w http.ResponseWriter, r *http.Request) {
	//This will be the 'hello world' for middlewares later
	fmt.Printf("Request received by %v\n", r.Host)
	message := fmt.Sprintf("%d clients connected", s.db.Clients)
	w.Write([]byte(message))
}

func (s *Server) initHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /hello", s.handleHello)
	mux.HandleFunc("GET /clients", s.handleClients)
	mux.HandleFunc("GET /messages", s.handleMessages)
	return mux
}

func initHandler(d *storage.Data) http.Handler {
	s := Server{
		db: d,
	}
	handler := s.initHandler()
	return handler
}
