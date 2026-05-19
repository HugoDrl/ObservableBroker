package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/HugoDrl/ObservableBroker.git/http"
	"github.com/HugoDrl/ObservableBroker.git/mqtt"
	"github.com/HugoDrl/ObservableBroker.git/storage"
)

func main() {
	//Will flag this kind of inputs
	filename := "logs/test.log"
	ttl := 10 * time.Minute
	d := make(chan(os.Signal), 1)
	db := storage.NewStorage(ttl)
	file, err := os.OpenFile(filename, os.O_CREATE, os.ModeAppend)
	if err != nil{
		fmt.Printf("File for logging is invalid [got: %s]\nWill log to stdout,\n[err: %s]\n", filename, err.Error())
		file = os.Stdout
	}
	logger := log.New(file, "", 0)

	mqtt.StartServer(db, logger)
	http.StartServer(db, logger)

	<-d
}
