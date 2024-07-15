package mongodb

import (
	"context"
	"errors"
	"github.com/cemayan/searchengine/constants"
	"github.com/cemayan/searchengine/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type MongoDB struct {
	client *mongo.Client
	ctx    context.Context
}

func (r MongoDB) Get(key string, params *[]string) (string, error) {
	defer func() {
		if err := r.client.Disconnect(r.ctx); err != nil {
			panic(err)
		}
	}()

	return "", nil
}

func (r MongoDB) Set(key string, value interface{}, params *[]string) error {
	defer func() {
		if err := r.client.Disconnect(r.ctx); err != nil {
			panic(err)
		}
	}()
	return nil
}

func New(projectName constants.Project) *MongoDB {
	cfg := config.GetConfig(projectName)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cli, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.Db.Persistent.Uri))

	if err != nil {
		panic(errors.New("an error occurred while connecting to the database"))
	}

	return &MongoDB{client: cli, ctx: ctx}
}
