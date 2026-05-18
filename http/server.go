package http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/HugoDrl/ObservableBroker.git/storage"
)

func initServer(d *storage.Data) *http.Server {
	h := initHandler(d)
	s := http.Server{Addr: ":8888", Handler: h}
	return &s
}

func StartServer(d *storage.Data) *http.Server {
	s := initServer(d)
	go s.ListenAndServe()
	// Don't want to do a custom format for ms, don't know is possible somehow else
	fmt.Printf("time=%s HTTP server started and listening on %s\n",time.Now().Format(time.RFC3339), s.Addr)
	return s
}
