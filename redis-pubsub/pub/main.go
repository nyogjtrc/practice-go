package main

import (
	"github.com/garyburd/redigo/redis"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewDevelopment()
	sugar := logger.Sugar()

	sugar.Info("redis pub")

	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		sugar.Fatal(err)
	}
	defer c.Close()

	c.Do("PUBLISH", "c1", "hello")
	c.Do("PUBLISH", "c1", "world")
	c.Do("PUBLISH", "c1", "new message to redis")

}
