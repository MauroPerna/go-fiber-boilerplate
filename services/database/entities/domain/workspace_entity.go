package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Workspace struct {
	ID           primitive.ObjectID   `bson:"_id,omitempty"`
	Owner        primitive.ObjectID   `bson:"owner"`
	Subscription primitive.ObjectID   `bson:"subscription,omitempty"`
	Companies    []primitive.ObjectID `bson:"companies"`
	CreatedAt    time.Time            `bson:"created_at"`
	UpdatedAt    time.Time            `bson:"updated_at"`
}
