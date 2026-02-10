package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Gender int

const (
	GenderUnspecified Gender = iota
	GenderMale
	GenderFemale
	GenderOther
)

type Profile struct {
	ID        string             `bson:"_id,omitempty" json:"id"`
	AccountID primitive.ObjectID `bson:"accountId" json:"accountId"`
	Username  string             `bson:"username" json:"username"`
	Avatar    string             `bson:"avatar" json:"avatar"`
	Fullname  string             `bson:"fullname" json:"fullname"`
	Bio       string             `bson:"bio" json:"bio"`
	Gender    Gender             `bson:"gender" json:"gender"`
}
