package messaging

import (
	"github.com/cemayan/searchengine/constants"
	"github.com/cemayan/searchengine/internal/config"
	"github.com/cemayan/searchengine/internal/messaging/nats"
	pb "github.com/cemayan/searchengine/protos/event"
	"github.com/nats-io/nats.go/jetstream"
)

var MessagingServer map[constants.Project]map[constants.Messaging]Messaging

func init() {
	messagingMap := make(map[constants.Project]map[constants.Messaging]Messaging)
	messagingMap[constants.Projection] = make(map[constants.Messaging]Messaging)
	messagingMap[constants.WriteApi] = make(map[constants.Messaging]Messaging)
	messagingMap[constants.ReadApi] = make(map[constants.Messaging]Messaging)
	MessagingServer = messagingMap
}

type Messaging interface {
	Publish(subj string, message *pb.Event) error
	PublishError(subj string, error *pb.SEError) error
	Subscribe(streamName string, consumerName string) jetstream.Consumer
	DeleteStream(name string)
}

func Init(projectName constants.Project) {
	m := config.GetConfig(projectName).Messaging

	if m.Nats != nil {

		MessagingServer[projectName][constants.Nats] = nats.New(projectName)

	} else if m.Kafka != nil {
		//TODO
	}
}
