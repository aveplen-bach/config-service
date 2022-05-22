package repo

import (
	"time"

	"github.com/aveplen-bach/config-service/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
)

func (r *Repository) GetFaceRecognitionConfig(ctx context.Context) (*model.FaceRecognitionConfig, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	collection := r.client.Database("configuration").Collection("facerec")

	count, err := collection.CountDocuments(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	if count == 0 {
		if _, err := collection.InsertOne(ctx, &model.FaceRecognitionConfig{
			ID:        primitive.NewObjectID(),
			Threshold: 0.6,
			CreatedAt: time.Now(),
		}); err != nil {
			return nil, err
		}
	}

	queryOptions := &options.FindOneOptions{}
	queryOptions.SetSort(bson.D{{"created_at", -1}})

	result := collection.FindOne(ctx, bson.D{}, queryOptions)
	if result.Err() != nil {
		return nil, err
	}

	var config model.FaceRecognitionConfig
	if err := result.Decode(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

func (r *Repository) UpdateFaceRecognitionConfig(ctx context.Context, config *model.FaceRecognitionConfig) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	
	collection := r.client.Database("configuration").Collection("facerec")

	_, err := collection.InsertOne(ctx, config)
	return err
}
