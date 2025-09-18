package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitMongoDB(uri string, timeout time.Duration) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	clientOptions := options.Client().ApplyURI(uri)
	db, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}
	// 测试连接
	if err = db.Ping(ctx, nil); err != nil {
		return nil, err
	}
	return db, nil
}
