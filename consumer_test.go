package kafkaavro_test

import (
	"testing"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	kafkaavro "github.com/mycujoo/go-kafka-avro"
)

func TestConsumer(t *testing.T) {

	kc, err := kafka.NewConsumer(&kafka.ConfigMap{
		"group.id":                 "gotest",
		"socket.timeout.ms":        1100,
		"session.timeout.ms":       10,
		"enable.auto.offset.store": false, // permit StoreOffsets()
	})
	if err != nil {
		t.Fatalf("%s", err)
	}

	srClient := &mockSchemaRegistryClient{}

	c, err := kafkaavro.NewConsumer([]string{"topic1"}, kc, srClient)
	if err != nil {
		t.Errorf("Subscribe failed: %s", err)
	}
	ch := make(chan struct{})
	c.Messages(ch)
	close(ch)
}

func TestConsumer_SubscribeTopics(t *testing.T) {
	kc, err := kafka.NewConsumer(&kafka.ConfigMap{
		"group.id":                 "gotest",
		"socket.timeout.ms":        1100,
		"session.timeout.ms":       10,
		"enable.auto.offset.store": false, // permit StoreOffsets()
	})
	if err != nil {
		t.Fatalf("%s", err)
	}

	srClient := &mockSchemaRegistryClient{}

	c, err := kafkaavro.NewConsumer(nil, kc, srClient)
	if err != nil {
		t.Errorf("Create failed: %s", err)
	}
	ch := make(chan struct{})
	c.Messages(ch)
	err = c.SubscribeTopics([]string{"topic1"}, nil)
	if err != nil {
		t.Errorf("Subscribe failed: %s", err)
	}

	close(ch)
}
