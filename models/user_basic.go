package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserBasic struct {
	Identity  primitive.ObjectID `bson:"_id" json:"_id"       `
	Account   string             `bson:"account" json:"account"   `
	Password  string             `bson:"password" json:"password"  `
	Nickname  string             `bson:"nickname" json:"nickname"  `
	Sex       int                `bson:"sex" json:"sex"       `
	Email     string             `bson:"email" json:"email"     `
	Avatar    string             `bson:"avatar" json:"avatar"    `
	CreatedAt int64              `bson:"created_at" json:"created_at"`
	UpdatedAt int64              `bson:"updated_at" json:"updated_at"`
}

func (userBasic *UserBasic) CollectionName() string {
	return "user_basic"
}

func GetUserBasicBy_AccountPassword(account, password string) (*UserBasic, error) {
	ub := new(UserBasic)
	err := Mongo.Collection((&UserBasic{}).CollectionName()).FindOne(context.Background(), bson.D{
		{Key: "account", Value: account},
		{Key: "password", Value: password},
	}).Decode(ub)
	//err := Mongo.Collection("user_basic").FindOne(context.Background(), bson.D{}).Decode(ub)
	return ub, err
}

func GetUserBasicBy_Identity(Identity primitive.ObjectID) (*UserBasic, error) {
	ub := new(UserBasic)
	err := Mongo.Collection((&UserBasic{}).CollectionName()).FindOne(context.Background(), bson.D{{Key: "_id", Value: Identity}}).Decode(ub)
	return ub, err
}

func GetUserBasicCountBy_Email(email string) (int64, error) {
	return Mongo.Collection((&UserBasic{}).CollectionName()).CountDocuments(context.Background(), bson.D{{Key: "email", Value: email}})
}

func InsertOneUserBasic(ub *UserBasic) error {
	_, err := Mongo.Collection((&UserBasic{}).CollectionName()).InsertOne(context.Background(), ub)
	return err
}
