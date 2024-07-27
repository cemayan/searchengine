package messaging

import (
	"github.com/cemayan/searchengine/constants"
	"github.com/cemayan/searchengine/internal/config"
	"github.com/cemayan/searchengine/internal/messaging/nats"
	pb "github.com/cemayan/searchengine/protos/event"
	"github.com/nats-io/nats.go/jetstream"
)

var MessagingServer Messaging

type Messaging interface {
	Publish(subj string, message *pb.Event) error
	Subscribe(streamName string, consumerName string) jetstream.Consumer
	DeleteStream(name string)
}

func Init(projectName constants.Project) {
	m := config.GetConfig(projectName).Messaging

	if m.Nats != nil {
		MessagingServer = nats.New(projectName)
	} else if m.Kafka != nil {
		//TODO
	}
}
