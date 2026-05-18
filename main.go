package main

import (
	"os"
	"time"

	"github.com/HugoDrl/ObservableBroker.git/http"
	"github.com/HugoDrl/ObservableBroker.git/mqtt"
	"github.com/HugoDrl/ObservableBroker.git/storage"
)

func main() {
	d := make(chan(os.Signal), 1)
	db := storage.NewStorage(10*time.Minute)

	mqtt.StartServer(db)
	http.StartServer(db)

	<-d
}
