package main

import (
	"github.com/garyburd/redigo/redis"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewDevelopment()
	sugar := logger.Sugar()

	sugar.Info("redis sub")

	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		sugar.Fatal(err)
	}
	defer c.Close()

	psc := redis.PubSubConn{Conn: c}

	err = psc.Subscribe("c1")
	if err != nil {
		sugar.Fatal(err)
	}

	done := make(chan error, 1)

	go func() {
		for {
			switch n := psc.Receive().(type) {
			case error:
				done <- err
			case redis.Message:
				sugar.Info(string(n.Data))
				done <- nil
			case redis.Subscription:
				sugar.Info(n)
				done <- nil
			}
		}
	}()

	for {
		err = <-done
		if err != nil {
			sugar.Warn(err)
		}
	}
}
