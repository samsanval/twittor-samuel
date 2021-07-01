package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoConnection = ConnectBD()
var clientOptions = options.Client().ApplyURI("mongodb://admin:example@localhost:2007")

func ConnectBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		return client
	}
	log.Println("Successful Connection")
	return client
}

/* CheckConnection */
func CheckConnection() int {
	err := MongoConnection.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		return 0
	}
	return 1
}
