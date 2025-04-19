package service

import (
	"context"
	"time"

	"go-api/internal/db"
	"go-api/internal/model"
)

func LogChargeRequest(log model.ChargeRequestLog) error {
	log.Timestamp = time.Now()

	_, err := db.ChargeLogs.InsertOne(context.Background(), log)
	return err
}
