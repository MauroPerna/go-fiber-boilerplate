package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Plan representa un documento en la colección "plans"
type Plan struct {
	ID                     primitive.ObjectID   `bson:"_id,omitempty"`
	Name                   string               `bson:"name"`
	Description            string               `bson:"description"`
	Price                  float64              `bson:"price"`
	Interval               string               `bson:"interval"`
	PlanType               string               `bson:"planType,omitempty"`
	MaxMembers             int                  `bson:"maxMembers,omitempty"`
	MaxOrdersPerMonth      int                  `bson:"maxOrdersPerMonth,omitempty"`
	MaxInvitationsPerMonth int                  `bson:"maxInvitationsPerMonth,omitempty"`
	MaxCompanies           int                  `bson:"maxCompanies,omitempty"`
	IsDeleted              bool                 `bson:"isDeleted"`
	StripePriceID          string               `bson:"stripePriceId"`
	CreatedBy              primitive.ObjectID   `bson:"createdBy,omitempty"`
	AssignedToWorkspaces   []primitive.ObjectID `bson:"assignedToWorkspaces,omitempty"`
	AssignedToRequests     []primitive.ObjectID `bson:"assignedToRequests,omitempty"`
	CreatedAt              time.Time            `bson:"created_at"`
	UpdatedAt              time.Time            `bson:"updated_at"`
}
