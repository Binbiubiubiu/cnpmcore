package util

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateObjectId() string {
	return primitive.NewObjectID().Hex()
}
