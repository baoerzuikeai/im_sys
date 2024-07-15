package models

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

type UserRoom struct {
	UserID string `bson:"user_identity"`
	RoomID string `bson:"room_identity"`
}

func (userroom *UserRoom) CollectionName() string {
	return "user_room"
}

func GetUsersByRoomIdentity(roomIdentity string) ([]string, error) {
	ur := new(UserRoom)
	users := make([]string, 0)
	cursor, err := Mongo.Collection((&UserRoom{}).CollectionName()).Find(context.Background(), bson.D{{
		Key:   "room_identity",
		Value: roomIdentity,
	}})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		if err = cursor.Decode(ur); err != nil {
			return nil, err
		}
		users = append(users, ur.UserID)
	}
	return users, nil
}
