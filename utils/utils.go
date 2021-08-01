package utils

import "go.mongodb.org/mongo-driver/mongo"

var Usercollection *mongo.Collection

func UserCollection(c *mongo.Database) {
	Usercollection = c.Collection("users")
}
