package db

import (
	"context"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

var DB *redis.Client
var ctx = context.Background()

var (
	HOST     = os.Getenv("REDIS_CACHE_HOST")
	PORT     = os.Getenv("REDIS_PORT")
	PASSWORD = os.Getenv("REDIS_CACHE_AUTH_TOKEN")
)

// connect creates a connection to redis
func Connect() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     HOST + ":" + PORT,
		Password: PASSWORD,
		DB:       0,
	})
	err := rdb.Ping(ctx).Err()
	if err != nil {
		logrus.Panic(err.Error())
	}
	logrus.Info("CONNECTED TO DB")
	DB = rdb
}
