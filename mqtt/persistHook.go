package mqtt

import (
	"bytes"
	"fmt"
	"time"

	"github.com/HugoDrl/ObservableBroker.git/storage"
	mqtt "github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/packets"
)

type persistHook struct {
	mqtt.HookBase
	Metrics *storage.Data
}

func NewPersistHook(storage *storage.Data) *persistHook{
	h := persistHook{Metrics: storage}
	return &h
}

func (h *persistHook) ID() string{
	return "persist-hook"
}

//Those will be the only hooks overwritten
func (h *persistHook) Provides(b byte) bool {
	return bytes.Contains([]byte{
		mqtt.OnConnect,
		mqtt.OnDisconnect,
		mqtt.OnPublish,
	}, []byte{b})
}

func (h *persistHook) OnConnect(cl *mqtt.Client, pk packets.Packet) error{
	fmt.Printf("new client connected ! [client: %s]\n", cl.ID)
	h.Metrics.Clients++
	return nil
}

func (h *persistHook) OnDisconnect(cl *mqtt.Client, err error, expire bool) {
	fmt.Printf("client disconnected ! [client: %s]\n", cl.ID)
	h.Metrics.Clients--
}

func (h *persistHook) OnPublish(cl *mqtt.Client, pk packets.Packet) (packets.Packet, error) {
	fmt.Printf("new message ! [client: %s message: %s]\n", cl.ID, pk.Payload)
	h.Metrics.Messages = append(h.Metrics.Messages, storage.Message{
		Time: time.Now(),
		Sender: cl.ID,
		Topic: pk.TopicName,
		Content: pk.Payload,
	})
	return pk, nil
}
