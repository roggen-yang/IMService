package server

import (
	"github.com/go-acme/lego/log"
	"github.com/micro/go-micro/broker"
)

type RabbitMqBroker struct {
	topic          string
	rabbitMqBroker broker.Broker
}

func NewRabbitMqBroker(topic string, rabbitMqBroker broker.Broker) (*RabbitMqBroker, error) {
	err := rabbitMqBroker.Init()
	if err != nil {
		return nil, err
	}

	err = rabbitMqBroker.Connect()
	if err != nil {
		return nil, err
	}

	return &RabbitMqBroker{topic: topic, rabbitMqBroker: rabbitMqBroker}, nil
}

func (r *RabbitMqBroker) Publisher(msg *broker.Message) {
	err := r.rabbitMqBroker.Publish(r.topic, msg)
	if err != nil {
		log.Printf("[publisher %s err]: %+v", r.topic, err)
	}
	log.Printf("[publisher %s]: %s", r.topic, string(msg.Body))
}

func (r *RabbitMqBroker) Subscribe(handlerFunc func(msg []byte) error) {
	_, err := r.rabbitMqBroker.Subscribe(r.topic, func(publication broker.Event) error {
		er := handlerFunc(publication.Message().Body)
		if er != nil {
			log.Println("handlerFunc msg err ", er)
		}
		return nil
	})
	if err != nil {
		log.Printf("[Subscribe %s err]: %+v", r.topic, err)
	}
	log.Printf("[publisher err]")
}
