package http

import (
	"fmt"
	"net/http"
	"time"
)

func initServer() *http.Server {
	h := initHandler()
	s := http.Server{Addr: ":8888", Handler: h}
	return &s
}

func StartServer() *http.Server {
	s := initServer()
	go s.ListenAndServe()
	// Don't want to do a custom format for ms, don't know is possible somehow else
	fmt.Printf("time=%s HTTP server started and listening on %s\n",time.Now().Format(time.RFC3339), s.Addr)
	return s
}
