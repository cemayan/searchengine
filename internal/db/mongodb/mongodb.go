package mongodb

import (
	"context"
	"errors"
	"github.com/cemayan/searchengine/constants"
	"github.com/cemayan/searchengine/internal/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.Background()

type MongoDB struct {
	client *mongo.Client
}

func (r MongoDB) Get(key string, params *[]string) (interface{}, error) {
	var mm map[string]interface{}
	err := r.client.Database(constants.MongoDbDatabase).Collection(key).FindOne(ctx, bson.M{}).Decode(&mm)

	if err != nil {
		if err.Error() != "document is nil" || err.Error() != "mongo: no documents in result" {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return mm, nil
}

func (r MongoDB) Set(key string, value interface{}, params *[]string) error {

	upsert := true
	update := bson.D{{"$set", value}}

	_, err := r.client.Database(constants.MongoDbDatabase).Collection(key).UpdateOne(ctx, bson.D{}, update,
		&options.UpdateOptions{Upsert: &upsert})
	if err != nil {
		return err
	}

	return nil
}

func New(projectName constants.Project) *MongoDB {
	cfg := config.GetConfig(projectName)

	cli, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.Db.Persistent.Uri))

	if err != nil {
		panic(errors.New("an error occurred while connecting to the database"))
	}

	return &MongoDB{client: cli}
}
