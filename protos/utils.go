package protos

import (
	pb "github.com/cemayan/searchengine/protos/event"
	"github.com/google/uuid"
	"time"
)

func GetEvent(data []byte, eventType pb.EventType) *pb.Event {
	now := time.Now()

	event := &pb.Event{
		Id:   uuid.NewString(),
		Type: eventType,
		Date: now.Unix(),
		Data: data,
	}

	return event
}
