package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

var (
	whereContent []bson.M
)


type Query struct{
	Cursor *mongo.Cursor
	SingleResult *mongo.SingleResult
	Collection *mongo.Collection
}

func NewQuery(collectionName string) Query {
	collection := mongoClient(collectionName)
	q := Query{Collection: collection}

	return q
}

func (q Query) Create(data interface{}) {
	if _, err := q.Collection.InsertOne(context.TODO(), data);err != nil {
		log.Println(err)
	}
}

func (q Query) Where(column string, value string) Query {
	queryContent := bson.M{column:value}
	whereContent = append(whereContent, queryContent)

	return q
}

func (q Query) WhereDate(column string, value time.Time, operator string) {
	var queryContent bson.M
	switch operator{
	case ">":
		queryContent = bson.M{column:bson.M{"$gt":value}}
		break
	case "<":
		queryContent = bson.M{column:bson.M{"$lt":value}}
		break
	case "=":
		queryContent = bson.M{column:bson.M{"$eq":value}}
		break
	case "<=":
		queryContent = bson.M{column:bson.M{"$lte":value}}
		break
	case ">=":
		queryContent = bson.M{column:bson.M{"$gte":value}}
		break
	default:
		queryContent = bson.M{column:bson.M{"$eq":value}}
		break
	}

	whereContent = append(whereContent, queryContent)
}

func (q Query) Update(id string, data interface{}) (interface{}, error) {
	err := q.Collection.FindOneAndUpdate(context.TODO(), whereContent, data).Err()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (q Query) Delete() error {
	return q.Collection.FindOneAndDelete(context.TODO(), whereContent).Err()
}

func (q Query) Get(result []interface{}) []interface{} {
	ctx := context.Background()
	c, err := q.Collection.Find(ctx, whereContent)
	if err != nil {
		log.Println(err)
	}
	c.All(ctx, result)

	return result
}

func (q Query) Paginate(page, limit int, result []interface{}) ([]interface{}, error) {
	ctx := context.Background()
    options := new(options.FindOptions)
    if limit != 0 {
        if page == 0 {
            page = 1
        }
        options.SetSkip(int64((page - 1) * limit))
        options.SetLimit(int64(limit))
    }
	c, err := q.Collection.Find(ctx, whereContent, options)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	
	c.All(ctx, result)

	return result, nil
}

func (q Query) First() (map[string]interface{}, error) {
	var data map[string]interface{}
	if err := q.Collection.FindOne(context.TODO(), whereContent).Decode(&data); err != nil {
		log.Println(err)
		return nil, err
	}

	return data, nil
}