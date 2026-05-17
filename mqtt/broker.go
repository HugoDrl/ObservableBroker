package mqtt

import (
	mqtt "github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/hooks/auth"
	"github.com/mochi-mqtt/server/v2/listeners"
)

func initServer() *mqtt.Server{
	server := mqtt.New(nil)
	server.AddHook(new(auth.AllowHook), nil)
	server.AddHook(new(persistHook), nil)
	tcp := listeners.NewTCP(listeners.Config{ID:"1", Address: ":1883"})
	server.AddListener(tcp)
	return server
}

func StartServer() {
	server := initServer()
	server.Serve()
}
