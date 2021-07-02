package bd

import (
	"context"
	"time"

	"github.com/samsanval/twittor-samuel/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ModifyRegister(user models.User, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConnection.Database("twittor-samuel")
	collection := db.Collection("users-copy")
	register := make(map[string]interface{})
	if len(user.Name) > 0 {
		register["name"] = user.Name
	}
	if len(user.Surname) > 0 {
		register["surname"] = user.Surname
	}
	if len(user.Banner) > 0 {
		register["banner"] = user.Banner
	}
	if len(user.Avatar) > 0 {
		register["avatar"] = user.Avatar
	}
	if len(user.Biography) > 0 {
		register["biography"] = user.Biography
	}
	if len(user.Location) > 0 {
		register["location"] = user.Location
	}
	if len(user.Website) > 0 {
		register["website"] = user.Website
	}
	register["datebirth"] = user.DateBirth
	updateString := bson.M{
		"$set": register,
	}
	objId, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{
		"_id": bson.M{"$eq": objId},
	}
	_, err := collection.UpdateOne(ctx, filter, updateString)
	if err != nil {
		return false, err
	}
	return true, nil
}
