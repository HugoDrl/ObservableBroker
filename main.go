package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/HugoDrl/ObservableBroker.git/http"
	"github.com/HugoDrl/ObservableBroker.git/mqtt"
	"github.com/HugoDrl/ObservableBroker.git/storage"
)

func main() {
	filename := flag.String("logfile", "", "log file to write to")
	ttlInt := flag.Int64("ttl", 0, "received messages' time to live (in seconds)")

	// Init storage with data time to live
	ttl := time.Duration(*ttlInt) * time.Minute
	db := storage.NewStorage(ttl)

	file, err := os.OpenFile(*filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.FileMode(int(0666)))
	if err != nil{
		fmt.Printf("Logging configuration is invalid [file: %s]\nWill log to stdout\n[err: %s]\n", *filename, err.Error())
		file = os.Stdout
	}
	logger := log.New(file, "INFO: ", log.Flags())

	mqtt.StartServer(db, logger)
	http.StartServer(db, logger)

	// Prevents program to stop, has to be manually stopped
	d := make(chan struct{})
	<-d
}
