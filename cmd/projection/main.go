package main

import (
	"context"
	"encoding/json"
	"flag"
	"github.com/cemayan/searchengine/constants"
	"github.com/cemayan/searchengine/internal/config"
	"github.com/cemayan/searchengine/internal/db"
	"github.com/cemayan/searchengine/internal/messaging"
	"github.com/cemayan/searchengine/internal/service"
	eventpb "github.com/cemayan/searchengine/protos/event"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"os/signal"
	"syscall"
)

var (
	configPath string // get scraper config yaml path as flag
)

func init() {
	flag.StringVar(&configPath, "config", "", "Path of config yaml")
	flag.Parse()
	// config initializer
	config.Init(constants.WriteApi, configPath)
	// db initializer
	db.Init(constants.WriteApi)
	// messaging initializer
	messaging.Init(constants.WriteApi)
}

func main() {
	cfg := config.GetConfig(constants.WriteApi)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	svc := service.NewWriteService(constants.WriteApi)

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(cfg.Db.Persistent.Uri), options.Client().SetReplicaSet("rs"))
	if err != nil {
		panic(err)
	}

	database := client.Database(constants.MongoDbDatabase)
	recordCollection := database.Collection(constants.DbName2Str[constants.Record])
	recordMetadataCollection := database.Collection(constants.DbName2Str[constants.RecordMetadata])

	changeRecordStream, err := recordCollection.Watch(context.TODO(), mongo.Pipeline{})
	if err != nil {
		panic(err)
	}

	changeRecordMetadataStream, err := recordMetadataCollection.Watch(context.TODO(), mongo.Pipeline{})
	if err != nil {
		panic(err)
	}

	go func() {
		for changeRecordStream.Next(context.TODO()) {
			var data bson.M
			if err := changeRecordStream.Decode(&data); err != nil {
				panic(err)
			}

			switch data["operationType"] {
			case "insert":
				rawData := data["fullDocument"].(primitive.M)
				if key, ok := rawData[constants.MongoDbRowKey].(string); ok {
					svc.PublishToNats([]byte(key), constants.NatsEventsStream, eventpb.EventType_RECORD_CREATED, eventpb.EntityType_Record)
				}
			}
		}
	}()

	go func() {
		for changeRecordMetadataStream.Next(context.TODO()) {
			var data bson.M
			if err := changeRecordMetadataStream.Decode(&data); err != nil {
				panic(err)
			}

			switch data["operationType"] {
			case "insert":
				rawData := data["fullDocument"].(primitive.M)
				marshal, _ := json.Marshal(rawData)
				svc.PublishToNats(marshal, constants.NatsEventsStream, eventpb.EventType_RECORDMETADATA_CREATED, eventpb.EntityType_RecordMetadata)
			}
		}
	}()

	<-done

}
