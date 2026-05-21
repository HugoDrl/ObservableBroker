package storage

import (
	"slices"
	"time"
)

type Message struct{
	Time time.Time
	Sender string
	Topic string
	Content []byte
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

type Data struct{
	Clients int16
	Messages []Message
	Topics map[string]int
}

func NewStorage(ttl time.Duration) *Data{
	d := Data{}
	d.Topics = make(map[string]int)
	go func(){
		for {
			d.cleanMessages(ttl)
			time.Sleep(10*time.Second)
		}
	}()
	return &d
}

func (d *Data) cleanMessages(from time.Duration) []Message {
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
	return deletedMessages
}
