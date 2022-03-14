package app

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"

	"ezpz/internals/database"

	"go.mongodb.org/mongo-driver/mongo"
)

const (
	UserCollection string = "users"
)

func Create(collectionName string, v interface{}) string {
	result, err := database.MongoClient(collectionName).InsertOne(context.TODO(), v)
	if err != nil {
		log.Println(err)
	}

	return result.InsertedID.(primitive.ObjectID).Hex()

}

func Find(CollectionName string, column string, value string) map[string]interface{} {
	var data map[string]interface{}
	if err := database.MongoClient(CollectionName).FindOne(context.TODO(), bson.D{{column, value}}).Decode(&data); err != nil {
		panic(err)
	}

	return data
}

func Get() {

}

func Update(collectionName string, id string, v interface{}) *mongo.UpdateResult {
	result, err := database.MongoClient(collectionName).UpdateByID(context.TODO(), id, v)
	if err != nil {
		log.Println(err)
	}

	return result
}

func Delete() {

}
