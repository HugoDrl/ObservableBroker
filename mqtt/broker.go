package mqtt

import (
	"github.com/HugoDrl/ObservableBroker.git/storage"
	mqtt "github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/hooks/auth"
	"github.com/mochi-mqtt/server/v2/listeners"
)

func initServer(hook mqtt.Hook) *mqtt.Server{
	server := mqtt.New(nil)
	server.AddHook(new(auth.AllowHook), nil)
	server.AddHook(hook, nil)
	tcp := listeners.NewTCP(listeners.Config{ID:"1", Address: ":1883"})
	server.AddListener(tcp)
	return server
}

func StartServer(storage *storage.Data) {
	hook := NewPersistHook(storage)
	server := initServer(hook)
	go server.Serve()
}
