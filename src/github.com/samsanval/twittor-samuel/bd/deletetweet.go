package bd

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteTweet(ID string, UserId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	database := MongoConnection.Database("twittor-samuel")
	collection := database.Collection("tweets")

	objId, _ := primitive.ObjectIDFromHex(ID)
	condition := bson.M{
		"_id":    objId,
		"userId": UserId,
	}
	_, err := collection.DeleteOne(ctx, condition)
	return err
}
