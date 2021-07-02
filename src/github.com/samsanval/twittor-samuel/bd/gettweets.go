package bd

import (
	"context"
	"log"
	"time"

	"github.com/samsanval/twittor-samuel/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetTweets(ID string, page int64) ([]*models.GetTweets, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	database := MongoConnection.Database("twittor-samuel")
	collection := database.Collection("tweets")
	var results []*models.GetTweets
	condition := bson.M{
		"userId": ID,
	}
	options := options.Find()
	options.SetLimit(20)
	options.SetSort(bson.D{{Key: "date", Value: -1}})
	options.SetSkip((page - 1) * 20)
	cursor, err := collection.Find(ctx, condition, options)
	if err != nil {
		log.Fatal(err.Error())
		return results, false
	}
	for cursor.Next(context.TODO()) {
		var register models.GetTweets
		err = cursor.Decode(&register)
		if err != nil {
			return results, false
		}
		results = append(results, &register)
	}
	return results, true

}
