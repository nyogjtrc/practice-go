package rank

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

type rank struct {
	Conn redis.Conn
	key  string
}

func NewRank(c redis.Conn) *rank {
	return &rank{
		Conn: c,
		key:  "rank",
	}
}

func (r *rank) NewScore(score int, member string) error {
	_, err := r.Conn.Do("ZADD", r.key, score, member)
	if err != nil {
		return fmt.Errorf("error when NewScore: %s", err)
	}
	return nil
}

func (r *rank) Top() {}
