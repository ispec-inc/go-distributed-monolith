package redis

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func New(c Config) (redis.Conn, error) {
	addr := fmt.Sprintf("%s:%s", c.Host, c.Port)

	conn, err := redis.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func NewPubSub(c PubSubConfig) (*redis.PubSubConn, error) {
	addr := fmt.Sprintf("%s:%s", c.Host, c.Port)

	conn, err := redis.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}

	return &redis.PubSubConn{Conn: conn}, nil
}
