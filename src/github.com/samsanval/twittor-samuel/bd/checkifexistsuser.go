package bd

import (
	"context"
	"time"

	"github.com/samsanval/twittor-samuel/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*Check if exists user*/
func CheckIfExistsUser(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoConnection.Database("twittor-samuel")
	collectionUser := db.Collection("users-copy")
	condition := bson.M{"email": email}
	var results models.User
	err := collectionUser.FindOne(ctx, condition).Decode(&results)
	ID := results.ID.Hex()
	if err != nil {
		return results, false, ID
	}
	return results, true, ID

}
