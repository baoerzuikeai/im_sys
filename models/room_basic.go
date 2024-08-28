package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RoomBasic struct {
	Identity      primitive.ObjectID `bson:"_id"`
	Number        string             `bson:"number"`
	User_identity string             `bson:"user_identity"`
}
