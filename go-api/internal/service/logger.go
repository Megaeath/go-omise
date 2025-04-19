package service

import (
	"context"
	"time"

	"go-api/internal/db"
	"go-api/internal/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func LogChargeRequest(log model.ChargeRequestLog) (primitive.ObjectID, error) {
	log.Timestamp = time.Now()

	result, err := db.ChargeLogs.InsertOne(context.Background(), log)
	if err != nil {
		return primitive.NilObjectID, err
	}

	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return primitive.NilObjectID, err
	}

	return insertedID, nil
}
