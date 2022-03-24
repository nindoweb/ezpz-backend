package models

import (
	"context"
	"log"

	"ezpz/pkg/database"

	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

const (
	USER_COLLECTION string = "users"
)

func Create(collectionName string, data interface{}) {
	if _, err := database.MongoClient(collectionName).InsertOne(context.TODO(), data);err != nil {
		log.Println(err)
	}
}

func Find(collectionName string, column string, value string) map[string]interface{} {
	var data map[string]interface{}
	if err := database.MongoClient(collectionName).FindOne(context.TODO(), bson.D{{column, value}}).Decode(&data); err != nil {
		panic(err)
	}

	return data
}

func Get() {

}

func Update(collectionName string, id string, data interface{}) *mongo.UpdateResult {
	result, err := database.MongoClient(collectionName).UpdateByID(context.TODO(), id, data)
	if err != nil {
		log.Println(err)
	}

	return result
}

func Delete() {

}
