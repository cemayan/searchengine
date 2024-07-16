package mongodb

import (
	"context"
	"errors"
	"github.com/cemayan/searchengine/constants"
	"github.com/cemayan/searchengine/internal/config"
	"github.com/cemayan/searchengine/internal/db/redis"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.Background()

type MongoDB struct {
	client *mongo.Client
	redis  *redis.Redis
}

func (r MongoDB) GetAll() interface{} {
	batchSize := int32(100)

	collections, err := r.client.Database(constants.RecordDatabase).ListCollectionNames(ctx, bson.D{}, &options.ListCollectionsOptions{
		BatchSize: &batchSize,
	})
	if err != nil {
		return nil
	}

	return collections
}

func (r MongoDB) Get(dbName constants.DbName, key string, params *[]string) (interface{}, error) {
	var mm map[string]interface{}
	err := r.client.Database(constants.DbName2Str[dbName]).Collection(key).FindOne(ctx, bson.M{}).Decode(&mm)

	if err != nil {
		if err.Error() != "document is nil" || err.Error() != "mongo: no documents in result" {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return mm, nil
}

func (r MongoDB) Set(dbName constants.DbName, key string, value interface{}, params *[]string) error {

	upsert := true
	update := bson.D{{"$set", value}}

	_, err := r.client.Database(constants.DbName2Str[dbName]).Collection(key).UpdateOne(ctx, bson.D{}, update,
		&options.UpdateOptions{Upsert: &upsert})
	if err != nil {
		return err
	}

	if r.redis != nil {
		go func() {
			err := r.redis.Set(dbName, key, value, params)
			if err != nil {
				logrus.Errorln("redis set err", err)
			}
		}()
	}

	return nil
}

func New(projectName constants.Project, redis *redis.Redis) *MongoDB {
	cfg := config.GetConfig(projectName)

	cli, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.Db.Persistent.Uri))

	if err != nil {
		panic(errors.New("an error occurred while connecting to the database"))
	}

	return &MongoDB{client: cli, redis: redis}
}
