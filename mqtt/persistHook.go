package mqtt

import (
	"bytes"
	"log"
	"time"

	"github.com/HugoDrl/ObservableBroker.git/storage"
	mqtt "github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/packets"
)

type persistHook struct {
	mqtt.HookBase
	storage storage.StorageWriter
	logger *log.Logger
}

func NewPersistHook(storage storage.StorageWriter, logger *log.Logger) *persistHook{
	h := persistHook{
		storage: storage,
		logger: logger,
	}
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
		mqtt.OnSubscribed,
	}, []byte{b})
}

func (h *persistHook) OnConnect(cl *mqtt.Client, pk packets.Packet) error{
	h.logger.Printf("new client connected ! [client: %s]\n", cl.ID)
	array := cl.State.Subscriptions.GetAll() 
	topics := make([]string, len(array))
	for topic := range array {
		topics = append(topics, topic)
	}
	event := storage.NewConnectionEvent(cl.ID, topics, true, time.Now())
	h.storage.InsertClientEvent(event, storage.CONNECTION)
	return nil
}

func (h *persistHook) OnDisconnect(cl *mqtt.Client, err error, expire bool) {
	h.logger.Printf("client disconnected ! [client: %s]\n", cl.ID)
	array := cl.State.Subscriptions.GetAll() 
	topics := make([]string, len(array))
	for topic := range array {
		topics = append(topics, topic)
	}
	event := storage.NewConnectionEvent(cl.ID, topics, false, time.Now())
	h.storage.InsertClientEvent(event, storage.DISCONNECTION)
}

func (h *persistHook) OnPublish(cl *mqtt.Client, pk packets.Packet) (packets.Packet, error) {
	h.logger.Printf("new message ! [client: %s message: %s]\n", cl.ID, pk.Payload)
	message := storage.Message{
		Time: time.Now(),
		Sender: cl.ID,
		Topic: pk.TopicName,
		Content: pk.Payload,
	}
	h.storage.InsertMessage(message)
	return pk, nil
}

func (h *persistHook) OnSubscribed(cl *mqtt.Client, pk packets.Packet, b []byte){
	h.logger.Printf("client subscribed ! [client: %s, topics: %v]\n", cl.ID, cl.State.Subscriptions.GetAll())
	array := cl.State.Subscriptions.GetAll()
	topics := make([]string, len(array))
	for topic := range array {
		topics = append(topics, topic)
	}
	h.storage.AddTopicsConnection(topics)
}

