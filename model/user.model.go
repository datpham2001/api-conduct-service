package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID              *primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	CreatedTime     *time.Time          `json:"createdTime,omitempty" bson:"created_time,omitempty"`
	LastUpdatedTime *time.Time          `json:"lastUpdatedTime,omitempty" bson:"last_updated_time,omitempty"`

	UserID       string  `json:"userId,omitempty" bson:"user_id,omitempty"`
	Username     string  `json:"username,omitempty" bson:"username,omitempty"`
	EmailAddress string  `json:"emailAddress,omitempty" bson:"email_address,omitempty"`
	Bio          string  `json:"bio,omitempty" bson:"bio,omitempty"`
	Avatar       *string `json:"avatar,omitempty" bson:"avatar,omitempty"`

	ComplexQuery []*bson.M `json:"-" bson:"$and,omitempty"`
}
