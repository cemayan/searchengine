package protos

import (
	pb "github.com/cemayan/searchengine/protos/event"
	"github.com/cemayan/searchengine/types"
	"github.com/google/uuid"
	"time"
)

func GetEvent(data []byte, eventType pb.EventType, entityType pb.EntityType) *pb.Event {
	now := time.Now()

	event := &pb.Event{
		Id:         uuid.NewString(),
		EntityType: entityType,
		Type:       eventType,
		Date:       now.Unix(),
		Data:       data,
	}

	return event
}

func GetError(err *types.SEError) *pb.SEError {
	now := time.Now()

	error := &pb.SEError{
		Kind:  string(err.Kind),
		Error: err.Error,
		Key:   err.Key,
		Value: err.Value,
		Date:  now.Unix(),
	}

	return error
}
