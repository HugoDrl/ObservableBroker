package http

import (
	"log"
	"net/http"

	"github.com/HugoDrl/ObservableBroker.git/storage"
)

func initServer(d *storage.Data, logger *log.Logger) *http.Server {
	h := initHandler(d, logger)
	s := http.Server{Addr: ":8888", Handler: h}
	return &s
}

func StartServer(d *storage.Data, logger *log.Logger) *http.Server {
	s := initServer(d, logger)
	go func() {
		logger.Fatal(s.ListenAndServe())
	}()

	logger.Printf("HTTP server started and listening on %s\n", s.Addr)
	return s
}
