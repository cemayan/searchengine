package redis

import (
	"context"
	"fmt"
	"github.com/cemayan/searchengine/constants"
	"github.com/cemayan/searchengine/internal/config"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

var ctx = context.Background()
var Client *redis.Client

// Redis represents of RedisDb Client
type Redis struct {
	client *redis.Client
}

func getClient(projectName constants.Project) *redis.Client {
	c := &redis.Client{}

	if Client != nil {
		c = Client
	} else {
		cfg := config.GetConfig(projectName)
		c = redis.NewClient(&redis.Options{
			Addr:     cfg.Db.Cache.Addr,
			Password: cfg.Db.Cache.Pass,
			DB:       0, // use default DB
		})
	}

	return c
}

func (r Redis) GetAll() interface{} {
	//TODO implement me
	panic("implement me")
}

// Get returns an  array according to given parameters
// In redis this method returns []interface{}
func (r Redis) Get(dbName constants.DbName, key string, params *[]string) (interface{}, error) {

	currentPath := params
	if params == nil {
		path := &[]string{}
		*path = append(*path, "$")
		// it means whole json
		currentPath = path
	}

	val, err := r.client.JSONGet(ctx, fmt.Sprintf("%s:%s", constants.DbName2Str[dbName], key), *currentPath...).Expanded()

	if err != nil {
		return "", err
	}

	return val, nil
}

// Set sets data to  redis with prefix
// for record : record, scraping results : recordmetadata
func (r Redis) Set(dbName constants.DbName, key string, value interface{}, params *[]string) error {

	currentPath := ""
	if params == nil {
		// it means whole json
		currentPath = "$"
	}

	_, err := r.client.JSONSet(ctx, fmt.Sprintf("%s:%s", constants.DbName2Str[dbName], key), currentPath, value).Result()
	if err != nil {
		return err
	}

	return nil
}

func New(projectName constants.Project) *Redis {

	r := &Redis{client: getClient(projectName)}

	if err := r.client.Ping(ctx).Err(); err != nil {
		logrus.Errorln(err)
	}

	return r
}
