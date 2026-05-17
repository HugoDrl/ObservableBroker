package main

import (
	"fmt"
	"time"

	"github.com/HugoDrl/ObservableBroker.git/mqtt"
	"github.com/HugoDrl/ObservableBroker.git/storage"
)

func main() {
	db := storage.NewStorage(10*time.Minute)
	mqtt.StartServer(db)
	for {
		time.Sleep(time.Second)
		fmt.Printf("%v\n", db)
	}
}
