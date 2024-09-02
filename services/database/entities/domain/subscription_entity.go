package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Subscription struct {
	ID                   primitive.ObjectID `bson:"_id,omitempty"`
	SubscriptionIDStripe string             `bson:"subscriptionIdStripe,omitempty"`
	Status               string             `bson:"status"`
	Plan                 primitive.ObjectID `bson:"plan"`
	CreatedAt            time.Time          `bson:"created_at"`
	UpdatedAt            time.Time          `bson:"updated_at"`
}
