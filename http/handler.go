package http

import (
	"fmt"
	"log"
	"net/http"

	"github.com/HugoDrl/ObservableBroker.git/storage"
)

type Server struct {
	db storage.StorageReader
	logger *log.Logger
}

func (s *Server) handleHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}

func (s *Server) handleMessages(w http.ResponseWriter, r *http.Request) {
	s.logger.Printf("Request received by %v\n", r.Host)
	messages := fmt.Sprintf("%v", s.db.ReadMessages())
	w.Write([]byte(messages))
}

func (s *Server) handleClients(w http.ResponseWriter, r *http.Request) {
	//This will be the 'hello world' for middlewares later
	s.logger.Printf("Request received by %v\n", r.Host)
	message := fmt.Sprintf("%d clients connected", s.db.GetConnectedClients())
	w.Write([]byte(message))
}

func (s *Server) handleClientEvents(w http.ResponseWriter, r *http.Request) {
	s.logger.Printf("Request received by %v\n", r.Host)
	message := fmt.Sprintf("%v", s.db.ReadEvents())
	w.Write([]byte(message))
}

func (s *Server) initHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /hello", s.handleHello)
	mux.HandleFunc("GET /clients", s.handleClients)
	mux.HandleFunc("GET /clients/events", s.handleClientEvents)
	mux.HandleFunc("GET /messages", s.handleMessages)
	return mux
}

func initHandler(d storage.StorageReader, logger *log.Logger) http.Handler {
	s := Server{
		db: d,
		logger: logger,
	}
	handler := s.initHandler()
	return handler
}
