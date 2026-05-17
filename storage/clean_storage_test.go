package storage

import (
	"testing"
	"time"
)

func TestDropOldStorage(t *testing.T) {
	d := Data{
		Clients: 0,
		Messages: []Message{{
			Time: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		}},
	}
	d.cleanMessages(0)
	expected := []Message{}
	if !MessageListEqual(d.Messages, expected) {
		t.Errorf("Error: Old storage badly dropped [expected: %v - got: %v]\n", expected, d.Messages)
	}
}

func TestZeroDate(t *testing.T) {
	d := Data{}
	d.cleanMessages(0)
	expected := []Message{}
	if !MessageListEqual(expected, d.Messages) {
		t.Errorf("Error: ZeroTime badly dropped [expected: %v - got: %v]\n", expected, d.Messages)
	}
}

func TestCleanRoutine(t *testing.T) {
	messageNow := Message{Time: time.Now()}
	messageOld := Message{Time: time.Now().Add(-time.Duration(100*time.Second))}

	d := Data{Messages: []Message{messageNow, messageOld}}
	d.cleanMessages(time.Duration(10*time.Second))

	expected := []Message{messageNow}
	if !MessageListEqual(expected, d.Messages) {
		t.Errorf("Error: Duration messages badly dropped [expected: %v - got: %v]\n", expected, d.Messages)
	}
}
