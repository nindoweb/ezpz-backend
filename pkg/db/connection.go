package db

import (
	"context"
	"fmt"
	"log"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func mongoClient(collection string) *mongo.Collection {
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
