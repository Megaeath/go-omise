package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoLogger struct {
	collection *mongo.Collection
}

func NewMongoLogger(ctx context.Context, uri, dbName, collectionName string) (*MongoLogger, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	coll := client.Database(dbName).Collection(collectionName)
	return &MongoLogger{collection: coll}, nil
}

func (l *MongoLogger) LogChargeProcess(ctx context.Context, logID, status string, amount int64, message string) error {
	doc := bson.M{
		"log_id":          logID,
		"status":          status,
		"amount_subunits": amount,
		"message":         message,
		"processed_at":    time.Now().UTC(),
	}

	_, err := l.collection.InsertOne(ctx, doc)
	return err
}

func (l *MongoLogger) GetLogByID(ctx context.Context, logID string) (bson.M, error) {
	var result bson.M
	err := l.collection.FindOne(ctx, bson.M{"log_id": logID}).Decode(&result)
	return result, err
}
