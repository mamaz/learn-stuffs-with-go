package consumer

import (
	"log"

	"github.com/nsqio/go-nsq"
)

type Handler struct{}

func (h *Handler) HandleMessage(m *nsq.Message) error {

	log.Println(string(m.Body))

	return nil
}

func Run() {

	config := nsq.NewConfig()
	consumer, err := nsq.NewConsumer("notification", "whatsapp", config)
	if err != nil {
		log.Fatalf("error on creating consumer %v", err)
	}

	consumer.AddHandler(&Handler{})

	consumer.ConnectToNSQLookupd("http://127.0.0.1:4160")
}
