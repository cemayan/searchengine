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

// MongoDB represents of MongoDb/Redis Client
type MongoDB struct {
	client      *mongo.Client
	redis       *redis.Redis
	cacheDirect bool
}

// GetAll returns whole collectionName
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

// Get returns an  array according to given parameters
// In redis this method returns map[string]interface{}
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

// Set sets data to  mongodb
// collections =>  record : record, scraping results : recordmetadata
// If mongodb save operation is successful this record will be saved to redis for caching
func (r MongoDB) Set(dbName constants.DbName, key string, value interface{}, params *[]string) error {

	upsert := true
	update := bson.D{{"$set", value}}

	_, err := r.client.Database(constants.DbName2Str[dbName]).Collection(key).UpdateOne(ctx, bson.D{}, update,
		&options.UpdateOptions{Upsert: &upsert})
	if err != nil {
		return err
	}

	// Write data to redis directly
	if r.cacheDirect {
		go func() {
			err := r.redis.Set(dbName, key, value, params)
			if err != nil {
				logrus.Errorln("redis set err", err)
			}
		}()
	}

	return nil
}

func New(projectName constants.Project) *MongoDB {
	cfg := config.GetConfig(projectName)

	r := &redis.Redis{}

	if cfg.CacheDirect.Enabled {
		r = redis.New(projectName)
	}

	cli, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.Db.Persistent.Uri))

	if err != nil {
		panic(errors.New("an error occurred while connecting to the database"))
	}

	return &MongoDB{client: cli, redis: r, cacheDirect: cfg.CacheDirect.Enabled}
}
