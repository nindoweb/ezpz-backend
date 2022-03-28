package tasks

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

const COLLECTION_NAME = "tasks"

func dbClient(collection string) *mongo.Collection {
	credential := options.Credential{
		Username: viper.GetString("DB_USER"),
		Password: viper.GetString("DB_PASSWORD"),
	}
	uri := fmt.Sprintf("mongodb://%s:%s", viper.GetString("DB_HOST"), viper.GetString("DB_PORT"))

	var clientOptions *options.ClientOptions
	if credential.Username == "" && credential.Password == "" {
		clientOptions = options.Client().ApplyURI(uri)
	} else {
		clientOptions = options.Client().ApplyURI(uri).SetAuth(credential)
	}

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Println(err)
	}

	if err = client.Ping(context.TODO(), nil); err != nil {
		log.Println(err)
	}
	
	return client.Database(viper.GetString("DB_NAME")).Collection(collection)
}

func create(data interface{}) error {
	if _, err := dbClient(COLLECTION_NAME).InsertOne(context.TODO(), data);err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func find(column string, value interface{}) map[string]interface{} {
	var data map[string]interface{}
	if err := dbClient(COLLECTION_NAME).FindOne(context.TODO(), bson.M{column: value}).Decode(&data); err != nil {
		log.Println(err)
	}

	return data
}

func avalibleQueues(name string, date time.Time) ([]Queue, error) {
	var cursor *mongo.Cursor
	var queues []Queue
	ctx := context.Background()

	cursor, err := dbClient(COLLECTION_NAME).Find(context.TODO(), bson.M{"name": name, "reserved_at": bson.M{"$lte": time.Now()}})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if err := cursor.All(ctx, queues); err != nil {
		log.Println(err)
	}

	return queues, nil
}