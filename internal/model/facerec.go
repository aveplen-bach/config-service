package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FaceRecognitionConfig struct {
	ID        primitive.ObjectID `bson:"_id"`
	Threshold float64            `bson:"threshold"`
	CreatedAt time.Time          `bson:"created_at"`
}
