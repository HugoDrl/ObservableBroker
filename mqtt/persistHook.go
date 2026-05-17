package mqtt

import (
	"bytes"
	"fmt"

	mqtt "github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/packets"
)

type persistHook struct {
	mqtt.HookBase
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
	return nil
}

func (h *persistHook) OnDisconnect(cl *mqtt.Client, err error, expire bool) {
	fmt.Printf("client disconnected ! [client: %s]\n", cl.ID)
}

func (h *persistHook) OnPublish(cl *mqtt.Client, pk packets.Packet) (packets.Packet, error) {
	fmt.Printf("new message ! [client: %s message: %s]\n", cl.ID, pk.Payload)
	return pk, nil
}
