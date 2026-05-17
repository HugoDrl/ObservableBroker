package main

import (
	"os"

	"github.com/HugoDrl/ObservableBroker.git/mqtt"
)

func main() {
	d := make(chan os.Signal, 1)
	go mqtt.StartServer()
	//Waits for interruption
	<- d
}
