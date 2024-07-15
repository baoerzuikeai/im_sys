package models


type UserStatus struct{
	Uid string `bson:"uid"`
	Status int `bson:"status"`
	Last_online int64 `bson:"last_online"`
}

func GetStatusbyUid()  {
	
}