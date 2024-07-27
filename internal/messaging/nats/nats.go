package nats

import (
	"context"
	"github.com/cemayan/searchengine/constants"
	"github.com/cemayan/searchengine/internal/config"
	pb "github.com/cemayan/searchengine/protos/event"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
)

var Client *nats.Conn
var JetStream jetstream.JetStream

type Nats struct {
	client    *nats.Conn
	jetstream jetstream.JetStream
	cfg       *config.Nats
	stream    jetstream.Stream
	consumer  jetstream.Consumer
}

func createStream(streamName string) jetstream.Stream {
	stream, err := JetStream.CreateOrUpdateStream(context.TODO(), jetstream.StreamConfig{
		Name: streamName,
	})
	if err != nil {
		logrus.Fatalf("an error occured while creating or updating nats stream: %v", err)
	}

	return stream
}

func createConsumer(streamName string, consumerName string) jetstream.Consumer {
	consumer, err := JetStream.CreateOrUpdateConsumer(context.TODO(), streamName, jetstream.ConsumerConfig{
		Name: consumerName,
	})
	if err != nil {
		logrus.Fatalf("an error occured while creating or updating nats consumer: %v", err)
	}
	return consumer
}

func (n *Nats) getStream(streamName string) jetstream.Stream {
	stream, err := n.jetstream.Stream(context.TODO(), streamName)
	if err != nil {
		logrus.Fatalf("an error occured while getting nats stream: %v", err)
	}

	return stream
}

func (n *Nats) getConsumer(streamName string, consumerName string) jetstream.Consumer {
	consumer, err := n.jetstream.Consumer(context.TODO(), streamName, consumerName)
	if err != nil {
		logrus.Fatalf("an error occured while getting nats consumer: %v", err)
	}

	return consumer
}

func (n *Nats) DeleteStream(name string) {
	n.jetstream.DeleteStream(context.TODO(), name)
}

func (n *Nats) Subscribe(streamName string, consumerName string) jetstream.Consumer {

	if n.cfg.IsJsEnabled {
		return n.getConsumer(streamName, consumerName)
	}
	return nil
}

func (n *Nats) Publish(subj string, message *pb.Event) error {

	messageBytes, _ := proto.Marshal(message)

	if n.cfg.IsJsEnabled {
		_, err := n.jetstream.Publish(context.TODO(), subj, messageBytes)
		if err != nil {
			return err
		}
	} else {
		err := n.client.Publish(subj, messageBytes)
		if err != nil {
			return err
		}
	}

	return nil
}

func New(projectName constants.Project) *Nats {

	m := config.GetConfig(projectName).Messaging.Nats

	nt := &Nats{}

	if Client == nil {
		nc, err := nats.Connect(m.Url)

		if err != nil {
			logrus.Fatal("an error occured while connecting to nats server")
		}

		js, err := jetstream.New(nc)

		if err != nil {
			logrus.Fatal("an error occured while connecting to jetstream")
		}

		nt.client = nc
		nt.jetstream = js

		JetStream = js

	} else {
		nt.client = Client
		nt.jetstream = JetStream
	}

	nt.cfg = m

	for _, v := range m.Streams {
		createStream(v)
	}

	for _, v := range m.Consumers {
		createConsumer(v.Stream, v.Name)
	}

	logrus.Infoln("nats client and jetstream  initialized")

	return nt
}
