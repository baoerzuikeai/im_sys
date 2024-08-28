package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PrivateMessageBasic struct {
	Identity            primitive.ObjectID `bson:"_id" json:"_id"`
	UserIdentity        string             `bson:"user_identity" json:"from_uid"`
	ReceiveUserIdentity string             `bson:"receive_user_identity" json:"receive_uid"`
	Data                string             `bson:"data" json:"msg"`
	CreatedAt           int64              `bson:"created_at" json:"created_at"`
	UpdatedAt           int64              `bson:"updated_at" json:"updated_at"`
}

type PublicMessageBasic struct {
	Identity      primitive.ObjectID `bson:"_id" json:"_id"`
	UserIdentity  string             `bson:"user_identity" json:"uid"`
	Room_identity string             `bson:"room_identity" json:"room_identity"`
	Data          string             `bson:"data" json:"msg"`
	CreatedAt     int64              `bson:"created_at" json:"created_at"`
	UpdatedAt     int64              `bson:"updated_at" json:"updated_at"`
}

func (privateMessageBasic *PrivateMessageBasic) CollectionName() string {
	return "private_message_basic"
}

func (publicMessageBasic *PublicMessageBasic) CollectionName() string {
	return "public_message_basic"
}

func InsertOnePrivateMsg(privatemsg PrivateMessageBasic) error {
	_, err := Mongo.Collection((&PrivateMessageBasic{}).CollectionName()).InsertOne(context.Background(), privatemsg)
	return err
}

func InsertOnePublicMsg(publivmsg PublicMessageBasic) error {
	_, err := Mongo.Collection((&PublicMessageBasic{}).CollectionName()).InsertOne(context.Background(), publivmsg)
	return err
}

func GetPublicMsgbyRooMidentity(roomidentity string, limit, skip *int64) ([]*PublicMessageBasic, error) {
	data := make([]*PublicMessageBasic, 0)
	cursor, err := Mongo.Collection((&PublicMessageBasic{}).CollectionName()).Find(context.Background(), bson.M{"room_identity": roomidentity},
		&options.FindOptions{
			Limit: limit,
			Skip:  skip,
			Sort:  bson.D{{Key: "created_at", Value: -1}},
		})
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.Background()) {
		pbmsg := new(PublicMessageBasic)
		err = cursor.Decode(pbmsg)
		if err != nil {
			return nil, err
		}
		data = append(data, pbmsg)
	}
	return data, nil
}


