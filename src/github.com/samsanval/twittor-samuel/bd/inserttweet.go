package bd

import (
	"context"
	"time"

	"github.com/samsanval/twittor-samuel/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertTweet(tweet models.NewTweet) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	database := MongoConnection.Database("twittor-samuel")
	collection := database.Collection("tweets")

	register := bson.M{
		"userId":  tweet.UserId,
		"message": tweet.Message,
		"date":    tweet.Date,
	}
	result, err := collection.InsertOne(ctx, register)
	if err != nil {
		return "", false, err
	}
	objId, _ := result.InsertedID.(primitive.ObjectID)
	return objId.String(), true, nil

}
