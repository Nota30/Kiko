package cache

import (
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

var (
	Ctx = context.Background()
	Client *redis.Client
)

func Connect() {
 	Client = redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "",
        DB:       0,
    })

	pong, err := Client.Ping(Ctx).Result()
	if err != nil {
		logrus.Fatal(err)
	} else {
		logrus.Info("Connected to Redis ", pong)
	}
}