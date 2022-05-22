package repo

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"sync"
)

type Repository struct {
	client *mongo.Client
	mu     *sync.Mutex
}

func NewRepository() (*Repository, error) {
	clientOptions := options.Client().ApplyURI("mongodb://root:example@localhost:27017/")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	if err := client.Ping(context.Background(), nil); err != nil {
		return nil, err
	}

	return &Repository{
		client: client,
		mu:     &sync.Mutex{},
	}, nil
}
