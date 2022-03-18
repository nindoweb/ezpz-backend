package database

import (
	"context"
	"ezpz/config"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoClient(collection string) *mongo.Collection {
	credential := options.Credential{
		Username: config.DBConfig()["username"],
		Password: config.DBConfig()["password"],
	}
	uri := fmt.Sprintf("mongodb://%s:%s", config.DBConfig()["host"], config.DBConfig()["port"])
	clientOptions := options.Client().ApplyURI(uri).SetAuth(credential)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Println(err)
	}

	return client.Database(config.DBConfig()["name"]).Collection(collection)
}
