package redis

import (
	"context"
	"fmt"
	"github.com/cemayan/searchengine/constants"
	"github.com/cemayan/searchengine/internal/config"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type Redis struct {
	client *redis.Client
}

func (r Redis) GetAll() interface{} {
	//TODO implement me
	panic("implement me")
}

func (r Redis) Get(key string, params *[]string) (interface{}, error) {

	currentPath := params
	if params == nil {
		path := &[]string{}
		*path = append(*path, "$")
		// it means whole json
		currentPath = path
	}

	val, err := r.client.JSONGet(ctx, fmt.Sprintf("%s:%s", constants.RedisJsonPrefix, key), *currentPath...).Expanded()

	if err != nil {
		return "", err
	}

	return val, nil
}

func (r Redis) Set(key string, value interface{}, params *[]string) error {

	currentPath := ""
	if params == nil {
		// it means whole json
		currentPath = "$"
	}

	_, err := r.client.JSONSet(ctx, fmt.Sprintf("%s:%s", constants.RedisJsonPrefix, key), currentPath, value).Result()
	if err != nil {
		return err
	}

	return nil
}

func New(projectName constants.Project) *Redis {
	cfg := config.GetConfig(projectName)
	rdb := redis.NewClient(&redis.Options{
		Addr: cfg.Db.Cache.Addr,
		DB:   0, // use default DB
	})
	return &Redis{client: rdb}
}
