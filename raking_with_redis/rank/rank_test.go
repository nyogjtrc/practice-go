package rank

import "testing"
import "github.com/garyburd/redigo/redis"

func setup() redis.Conn {
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		panic(err)
	}

	return c
}

func teardown(c redis.Conn) {
	c.Do("DEL", "rank")
	c.Close()
}

func TestAdd(t *testing.T) {
	c := setup()
	defer teardown(c)

	rk := NewRank(c)
	rk.NewScore(1, "aaa")
	rk.NewScore(2, "bbb")
	rk.NewScore(3, "ccc")

	result, _ := redis.StringMap(c.Do("ZRANGE", "rank", 0, -1, "WITHSCORES"))
	t.Log(result)
}
