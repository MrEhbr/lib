package mqtt

import (
	"errors"

	mq "github.com/eclipse/paho.mqtt.golang"
)

// Publisher publishes messages on a transport
type Publisher interface {
	Publish(destination string, message []byte, qos int, persist bool) error
}

// Subscriber receives messages from a transport
type Subscriber interface {
	Subscribe(source string, qos int, callback func(Message))
	SubscribeMultiple(filters map[string]byte, callback func(Message))
	Unsubscribe(sources ...string)
}

// PublishSubscriber can both publish and receives messages from a transport
type PublishSubscriber interface {
	Publisher
	Subscriber
}

type Message interface {
	Topic() string
	Payload() []byte
}

type mqttMessenger struct {
	client Client
}

func NewPublishSubscriber(client Client) PublishSubscriber {
	return &mqttMessenger{
		client: client,
	}
}

func (m *mqttMessenger) Publish(topic string, msg []byte, qos int, retain bool) error {
	if topic == "" {
		return errors.New("empty topic")
	}

	if token := m.client.Publish(topic, byte(qos), retain, msg); token.Wait() && token.Error() != nil {
		return token.Error()
	}

	return nil
}

func (m *mqttMessenger) Subscribe(topic string, qos int, callback func(Message)) {
	m.client.Subscribe(topic, byte(qos), func(c mq.Client, msg mq.Message) {
		callback(msg)
	})
}

func (m *mqttMessenger) SubscribeMultiple(subs map[string]byte, callback func(Message)) {
	m.client.SubscribeMultiple(subs, func(c mq.Client, msg mq.Message) {
		callback(msg)
	})
}

func (m *mqttMessenger) Unsubscribe(topics ...string) {
	m.client.Unsubscribe(topics...)
}
