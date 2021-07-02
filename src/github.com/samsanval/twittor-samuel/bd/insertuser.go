package bd

import (
	"context"
	"time"

	"github.com/samsanval/twittor-samuel/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*Insert register*/
func InsertUser(user models.User) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoConnection.Database("twittor-samuel")
	collectionUsers := db.Collection("users-copy")
	user.Password, _ = EncryptPassword(user.Password)
	result, err := collectionUsers.InsertOne(ctx, user)
	if err != nil {
		return "", false, err
	}
	ObjId, _ := result.InsertedID.(primitive.ObjectID)
	return ObjId.String(), true, nil
}
