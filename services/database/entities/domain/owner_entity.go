package entities

import (
	"time"

	"github.com/MauroPerna/go-fiber-backend-test/common/enums"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Owner struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	UID            string             `bson:"uid"`
	ChatID         string             `bson:"chatId"`
	Email          string             `bson:"email"`
	Status         enums.Status       `bson:"status"`
	FirstName      string             `bson:"firstName,omitempty"`
	LastName       string             `bson:"lastName,omitempty"`
	Phone          string             `bson:"phone,omitempty"`
	ProfilePicture string             `bson:"profilePicture,omitempty"`
	Company        primitive.ObjectID `bson:"company,omitempty"`
	Workspace      Workspace          `bson:"workspace"`
	Language       enums.Language     `bson:"language"`
	LastConnection time.Time          `bson:"lastConnection,omitempty"`
	CreatedAt      time.Time          `bson:"created_at"`
	UpdatedAt      time.Time          `bson:"updated_at"`
}
