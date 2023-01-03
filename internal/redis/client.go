package redis

import (
	"github.com/cro4k/authorize/config"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

var c *redis.Client

func init() {
	client := redis.NewClient(&redis.Options{
		Addr:     config.C().Redis.Addr(),
		Password: config.C().Redis.Pass,
	})
	if err := client.Ping().Err(); err != nil {
		logrus.Fatal(err)
	}
	c = client
}

func CLI() *redis.Client {
	return c
}
