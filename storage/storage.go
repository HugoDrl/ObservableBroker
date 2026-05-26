package storage

import (
	"fmt"
	"slices"
	"time"
	"errors"
)

type Message struct{
	Time time.Time
	Sender string
	Topic string
	Content []byte
}

type Data struct{
	ClientsConnected int32
	Messages []Message
	Topics map[string]int
	ClientEvents []ClientEvent
}

type StorageWriter interface {
	InsertMessage(message Message)error
	InsertClientEvent(event Event, eventType EventType)error
	AddTopicsConnection(topics []string)error
}

func (m *Message) Equal(m2 *Message) bool {
	if m.Time != m2.Time {return false}
	if m.Sender != m2.Sender {return false}
	if m.Topic != m2.Topic {return false}
	if !slices.Equal(m.Content, m2.Content) {return false}
	return true
}

func MessageListEqual(l1, l2 []Message) bool {
	if len(l1) != len(l2) {return false}
	for i := range l1 {
		if !l1[i].Equal(&l2[i]) {return false}
	}
	return true
}

func NewStorage(ttl time.Duration) *Data{
	d := Data{}
	d.Topics = make(map[string]int)
	go func(){
		if ttl == 0 {
			return
		}
		for {
			d.cleanMessages(ttl)
			time.Sleep(10*time.Second)
		}
	}()
	return &d
}

func (d *Data) InsertMessage(message Message)error {
	// Maybe add errors later, like full messages queue or things like that
	d.Messages = append(d.Messages, message)
	return nil
}

func (d *Data) AddTopicsConnection(topics []string)error {
	for _, topic := range topics {
		d.Topics[topic]++
	}
	return nil
}

func (d *Data) InsertClientEvent(event Event, eventType EventType)error {
	if eventType != CONNECTION && eventType != DISCONNECTION {
		return fmt.Errorf("wrong connection type [connection type: %v]", eventType)
	}
	v, ok := event.(*ConnectionEvent)
	if !ok {
		return errors.New("only connection events are supported for now")
	}

	e := NewClientEvent(v.ClientId, eventType, time.Now())
	d.ClientEvents = append(d.ClientEvents, *e)

	if eventType == CONNECTION {
		d.connect(v.topics)
		return nil
	}

	//It can only be disconnect (for now)
	return d.disconnect(v.topics)
}

func (d *Data) connect(topics []string) {
	d.ClientsConnected++
	d.subscribe(topics)
}

func (d *Data) disconnect(topics []string) error {
	if err := d.unsubscribe(topics); err != nil {
		return err
	}

	if d.ClientsConnected == 0 {
		return errors.New("disconnection of client would have set number of clients to negative")
	}

	d.ClientsConnected--
	return nil
}

func (d *Data) subscribe(topics []string){
	for _, topic := range topics {
		d.Topics[topic]++
	}
}

func (d *Data) unsubscribe(topics []string) error {
	for _, topic := range topics {
		if t, ok := d.Topics[topic]; t == 0 || !ok{
			return fmt.Errorf("topic %s is already empty, unable to unsubscribe", topic)
		}
		d.Topics[topic]--
	}
		return nil
}

func (d *Data) cleanMessages(from time.Duration) ([]Message, error) {
	if from < 0 {
		return nil, fmt.Errorf("from can not be negative. [from: %d]", from)
	}
	startingDate := time.Now().Add(-from)
	deletedMessages := []Message{}
	i := 0
	for {
		if i >= len(d.Messages) {
			break
		}
		if time.Time.Compare(d.Messages[i].Time, startingDate) < 0 {
			deletedMessages = append(deletedMessages, d.Messages[i])
			d.Messages[i] = d.Messages[len(d.Messages)-1]
			d.Messages = d.Messages[:len(d.Messages)-1]
		}
		i++
	}
	return deletedMessages, nil
}
